/**
 * Represents the result of an API operation.
 */
export class Result {
  data: Record<string, any>;
  meta: Record<string, any>;
  ok: boolean;
  error?: string | null;
  statusCode: number;

  /**
   * Construct a Result object.
   *
   * @param props.data - Response data from the API.
   * @param props.meta - Metadata associated with the response.
   * @param props.ok - Whether the operation was successful.
   * @param props.error - Error message if operation failed.
   * @param props.statusCode - HTTP status code of the response.
   */
  constructor({
    data = {},
    meta = {},
    ok = true,
    error = null,
    statusCode = 200,
  }: {
    data?: Record<string, any>;
    meta?: Record<string, any>;
    ok?: boolean;
    error?: string | null;
    statusCode?: number;
  }) {
    this.data = data;
    this.meta = meta;
    this.ok = ok;
    this.error = error;
    this.statusCode = statusCode;
  }

  /**
   * Convert the Result object to a plain object.
   */
  toObject() {
    return {
      data: this.data,
      meta: this.meta,
      ok: this.ok,
      error: this.error,
      statusCode: this.statusCode,
    };
  }

  /**
   * String representation of the Result.
   */
  toString() {
    return `<Result ok=${this.ok} status=${this.statusCode} error=${this.error}>`;
  }
}

/**
 * Represents a successful API operation result.
 */
export class ResultOK extends Result {
  constructor(props: any = {}) {
    super({ ...props, ok: true });
  }
}

/**
 * Represents a failed API operation result.
 */
export class ResultError extends Result {
  constructor(props: any = {}) {
    super({ ...props, ok: false });
  }
}
