import { Result, ResultOK, ResultError } from "../src/result";
import { Client, SinglebaseError } from "../src/client";
import { uploadPresignedFile } from "../src/upload";

jest.mock("fs", () => ({
  createReadStream: jest.fn().mockReturnValue("mock-stream"),
}));

global.fetch = jest.fn();

// ---------------------------
// Result Tests
// ---------------------------

test("Result toObject and toString", () => {
  const r = new Result({ data: { foo: "bar" }, statusCode: 201 });
  const obj = r.toObject();
  expect(obj.data.foo).toBe("bar");
  expect(r.toString()).toContain("Result");
});

test("ResultOK and ResultError", () => {
  const ok = new ResultOK({ data: { success: true } });
  const err = new ResultError({ error: "failed", statusCode: 400 });
  expect(ok.ok).toBe(true);
  expect(err.ok).toBe(false);
  expect(err.statusCode).toBe(400);
});

// ---------------------------
// Client Tests
// ---------------------------

test("Client.validatePayload valid", () => {
  const client = new Client({ apiKey: "abc", endpointKey: "test" });
  const payload = { op: "create", foo: 123 };
  expect(client.validatePayload(payload)).toEqual(payload);
});

test("Client.validatePayload invalid", () => {
  const client = new Client({ apiKey: "abc", endpointKey: "test" });
  expect(() => client.validatePayload({} as any)).toThrow(SinglebaseError);
  expect(() => client.validatePayload("bad" as any)).toThrow(SinglebaseError);
});

test("Client.call success", async () => {
  (global.fetch as jest.Mock).mockResolvedValueOnce({
    ok: true,
    status: 200,
    json: async () => ({ data: { msg: "ok" }, meta: { page: 1 } }),
  });

  const client = new Client({ apiKey: "abc", endpointKey: "test" });
  const result = await client.call({ op: "ping" });
  expect(result).toBeInstanceOf(ResultOK);
  expect(result.data.msg).toBe("ok");
});

test("Client.call error", async () => {
  (global.fetch as jest.Mock).mockResolvedValueOnce({
    ok: false,
    status: 400,
    json: async () => ({ error: "Bad Request" }),
  });

  const client = new Client({ apiKey: "abc", endpointKey: "test" });
  const result = await client.call({ op: "ping" });
  expect(result).toBeInstanceOf(ResultError);
  expect(result.error).toBe("Bad Request");
});

// ---------------------------
// File Upload Tests
// ---------------------------

test("uploadPresignedFile success", async () => {
  (global.fetch as jest.Mock).mockResolvedValueOnce({
    ok: true,
    status: 204,
  });

  const result = await uploadPresignedFile("test.txt", {
    url: "http://fake-url",
    fields: {},
  });
  expect(result).toBe(true);
});

test("uploadPresignedFile failure", async () => {
  (global.fetch as jest.Mock).mockResolvedValueOnce({
    ok: false,
    status: 403,
  });

  await expect(
    uploadPresignedFile("test.txt", { url: "http://fake-url", fields: {} })
  ).rejects.toThrow("Upload failed: 403");
});
