export function getStorage(key: string): string | null {
  return localStorage.getItem(key)
}
