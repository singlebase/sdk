import * as fs from "fs";
import FormData from "form-data";

/**
 * Upload a file to a presigned URL.
 *
 * @param filepath - Path to the file to upload.
 * @param data - Presigned URL data.
 *   - `url`: Upload URL.
 *   - `fields`: Form fields required by the server.
 *
 * @returns {Promise<boolean>} True if the upload succeeded, otherwise throws.
 *
 * @throws {Error} If the upload fails with a non-2xx status.
 */
export async function uploadPresignedFile(
  filepath: string,
  data: { url: string; fields: Record<string, string> }
): Promise<boolean> {
  const form = new FormData();
  for (const [key, value] of Object.entries(data.fields)) {
    form.append(key, value);
  }
  form.append("file", fs.createReadStream(filepath));

  const resp = await fetch(data.url, {
    method: "POST",
    body: form as any,
    headers: form.getHeaders(),
  });

  if (resp.ok) return true;
  throw new Error(`Upload failed: ${resp.status}`);
}
