import { z } from 'zod';

export const loginSchema = z.object({
  email: z
    .string()
    .min(1, { message: "Email has to be filled" })
    .email("Email is not valid"),
  password: z
    .string()
    .min(6, { message: "Password has to be filled and include minimum 6 characters" })
});

export type LoginFormData = z.infer<typeof loginSchema>; 