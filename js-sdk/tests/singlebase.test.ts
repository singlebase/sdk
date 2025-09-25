import { Result, ResultOK, ResultError } from "../src/result";
import { Client, SinglebaseError } from "../src/client";

global.fetch = jest.fn();

// ---------------------------
// Result Tests
// ---------------------------

test("Result getData success and default", () => {
  const r = new ResultOK({
    data: {
      address: {
        city: {
          city_fullname: "San Francisco",
          zipcode: 94107,
        },
      },
    },
  });

  // Full data
  expect(r.getData()).toEqual(r.data);

  // Nested dot path
  expect(r.getData("address.city.city_fullname")).toBe("San Francisco");
  expect(r.getData("address.city.zipcode")).toBe(94107);

  // Missing key returns default
  expect(r.getData("address.country", "USA")).toBe("USA");
});

test("Result getData throws TypeError when wrong type", () => {
  const r = new ResultOK({ data: { user: { id: 123 } } });
  expect(() => r.getData("user.id.value")).toThrow(TypeError);
});

// ---------------------------
// Client Tests
// ---------------------------

test("Client.dispatch success", async () => {
  (global.fetch as jest.Mock).mockResolvedValueOnce({
    ok: true,
    status: 200,
    json: async () => ({ data: { msg: "ok" }, meta: {} }),
  });

  const client = new Client({ apiKey: "abc", endpointKey: "test" });
  const result = await client.dispatch({ op: "ping" });
  expect(result).toBeInstanceOf(ResultOK);
  expect(result.data.msg).toBe("ok");
});

test("Client.dispatch error", async () => {
  (global.fetch as jest.Mock).mockResolvedValueOnce({
    ok: false,
    status: 400,
    json: async () => ({ error: "Bad Request" }),
  });

  const client = new Client({ apiKey: "abc", endpointKey: "test" });
  const result = await client.dispatch({ op: "ping" });
  expect(result).toBeInstanceOf(ResultError);
  expect(result.error).toBe("Bad Request");
});

test("Client.validatePayload throws error", () => {
  const client = new Client({ apiKey: "abc", endpointKey: "test" });
  expect(() => client.validatePayload({} as any)).toThrow(SinglebaseError);
});
