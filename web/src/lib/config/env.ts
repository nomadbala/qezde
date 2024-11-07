const requiredEnvVars = [
  'NEXT_PUBLIC_API_URL',
  'NEXT_PUBLIC_SUPERTOKENS_APP_NAME',
] as const;

export const env = {
  apiUrl: process.env.NEXT_PUBLIC_API_URL,
  supertokensAppName: process.env.NEXT_PUBLIC_SUPERTOKENS_APP_NAME,
} as const;

// Validate environment variables at build time
requiredEnvVars.forEach((envVar) => {
  if (!process.env[envVar]) {
    throw new Error(`Missing required environment variable: ${envVar}`);
  }
}); 