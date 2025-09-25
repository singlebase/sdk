export class Result {
  data: Record<string, any>;
  meta: Record<string, any>;
  ok: boolean;
  error?: string | null;
  statusCode: number;

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

  /**
   * Retrieve a value from `data` using dot notation.
   * If no path is given, returns full data object.
   * If path not found, returns defaultValue.
   * Throws TypeError if traversal encounters non-object value.
   */
  getData(path?: string, defaultValue: any = null): any {
    if (!path) return this.data;

    let current: any = this.data;
    for (const part of path.split(".")) {
      if (typeof current !== "object" || current === null) {
        throw new TypeError(
          `Cannot traverse '${part}' — expected object, got ${typeof current}`
        );
      }
      if (!(part in current)) {
        return defaultValue;
      }
      current = current[part];
    }
    return current;
  }
}

export class ResultOK extends Result {
  constructor(props: any = {}) {
    super({ ...props, ok: true });
  }
}

export class ResultError extends Result {
  constructor(props: any = {}) {
    super({ ...props, ok: false });
  }
}
