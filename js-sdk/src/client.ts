import { Result, ResultOK, ResultError } from "./result";

/**
 * Type definition for API payloads.
 * Must contain 'op' field but allows arbitrary extra fields.
 */
export interface PayloadType {
  op: string;
  [key: string]: any;
}

/**
 * Base error type for Singlebase SDK.
 */
export class SinglebaseError extends Error {}

/**
 * Singlebase API client for making synchronous and asynchronous API calls.
 */
export class Client {
  private static BASE_API_URL = "https://cloud.singlebaseapis.com/api";
  private apiKey: string;
  private apiUrl: string;
  private headers: Record<string, string>;

  /**
   * Construct a new Client.
   *
   * @param apiKey - Your Singlebase API key.
   * @param apiUrl - Full API URL (optional). If not provided, `endpointKey` must be given.
   * @param endpointKey - Endpoint identifier appended to the base URL.
   * @param headers - Default headers for all requests.
   *
   * @throws {SinglebaseError} If `apiKey` is missing or no endpoint is defined.
   */
  constructor({
    apiKey,
    apiUrl,
    endpointKey,
    headers,
  }: {
    apiKey: string;
    apiUrl?: string;
    endpointKey?: string;
    headers?: Record<string, string>;
  }) {
    if (!apiKey) throw new SinglebaseError("MISSING_API_KEY");
    if (!apiUrl && !endpointKey)
      throw new SinglebaseError("MISSING_ENDPOINT_KEY");

    this.apiKey = apiKey;
    this.headers = headers || {};
    this.apiUrl = apiUrl || `${Client.BASE_API_URL}/${endpointKey}`;
  }

  /**
   * Validate that the payload contains a valid 'op' field.
   *
   * @param payload - The payload to validate.
   * @throws {SinglebaseError} If payload is missing or invalid.
   */
  validatePayload(payload: PayloadType) {
    if (typeof payload !== "object" || !payload.op) {
      throw new SinglebaseError("INVALID_PAYLOAD: missing 'op'");
    }
    return payload;
  }

  /**
   * Call the Singlebase API with the given payload.
   *
   * @param payload - The request payload (must include 'op').
   * @param headers - Optional additional headers to merge with defaults.
   * @param bearerToken - Optional bearer token for Authorization.
   *
   * @returns {Promise<Result>} The API response wrapped in a Result object.
   */
  async call(
    payload: PayloadType,
    headers?: Record<string, string>,
    bearerToken?: string
  ): Promise<Result> {
    try {
      this.validatePayload(payload);

      const _headers: Record<string, string> = {
        ...this.headers,
        ...(headers || {}),
        "x-api-key": this.apiKey,
        "x-sbc-sdk-client": "singlebase-js",
        "Content-Type": "application/json",
      };
      if (bearerToken) {
        _headers["Authorization"] = `Bearer ${bearerToken}`;
      }

      const resp = await fetch(this.apiUrl, {
        method: "POST",
        headers: _headers,
        body: JSON.stringify(payload),
      });

      const parsed = await resp.json();

      if (resp.ok) {
        return new ResultOK({
          data: parsed.data,
          meta: parsed.meta,
          statusCode: resp.status,
        });
      } else {
        return new ResultError({
          error: parsed.error,
          statusCode: resp.status,
        });
      }
    } catch (err: any) {
      return new ResultError({
        error: `EXCEPTION: ${err.message}`,
        statusCode: 500,
      });
    }
  }
}
