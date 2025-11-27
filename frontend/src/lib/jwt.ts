/**
 * JWT utility functions for token management
 */

interface JWTPayload {
  exp: number; // Expiration time (seconds since epoch)
  iat: number; // Issued at time
  sub: string; // Subject (user ID)
  [key: string]: unknown;
}

/**
 * Decode JWT token without verification
 * Note: This is only for reading the payload, not for security validation
 */
export function decodeJWT(token: string): JWTPayload | null {
  try {
    const parts = token.split(".");
    if (parts.length !== 3) {
      return null;
    }

    const payload = parts[1];
    const decoded = atob(payload.replace(/-/g, "+").replace(/_/g, "/"));
    return JSON.parse(decoded) as JWTPayload;
  } catch (error) {
    console.error("Failed to decode JWT:", error);
    return null;
  }
}

/**
 * Check if JWT token is expired
 * @param token - JWT token string
 * @param bufferSeconds - Optional buffer time in seconds (default: 0)
 * @returns true if token is expired or will expire within buffer time
 */
export function isTokenExpired(
  token: string,
  bufferSeconds: number = 0
): boolean {
  const payload = decodeJWT(token);
  if (!payload || !payload.exp) {
    return true;
  }

  const currentTime = Math.floor(Date.now() / 1000);
  return payload.exp - bufferSeconds <= currentTime;
}

/**
 * Get token expiration time in seconds
 * @param token - JWT token string
 * @returns expiration time in seconds, or null if invalid
 */
export function getTokenExpiration(token: string): number | null {
  const payload = decodeJWT(token);
  return payload?.exp ?? null;
}

/**
 * Get time until token expires in seconds
 * @param token - JWT token string
 * @returns seconds until expiration, or 0 if expired/invalid
 */
export function getTimeUntilExpiration(token: string): number {
  const exp = getTokenExpiration(token);
  if (!exp) {
    return 0;
  }

  const currentTime = Math.floor(Date.now() / 1000);
  const timeLeft = exp - currentTime;
  return Math.max(0, timeLeft);
}
