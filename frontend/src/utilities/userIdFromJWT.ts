function base64UrlDecode(str: string) {
    const padded = str.padEnd(str.length + ((4 - (str.length % 4)) % 4), "=");
    const base64 = padded.replace(/-/g, "+").replace(/_/g, "/");
    return atob(base64);
}

export function decodeJWT(token: string) {
    if (!token) return { header: "", payload: "", signature: "" };
    const [h, p, s] = token.split(".");
    const header = JSON.parse(base64UrlDecode(h));
    const payload = JSON.parse(base64UrlDecode(p));
    return { header, payload, signature: s };
}
