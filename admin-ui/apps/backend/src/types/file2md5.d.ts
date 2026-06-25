declare module 'file2md5' {
  export default function file2md5(file: Blob): Promise<string>;
}
