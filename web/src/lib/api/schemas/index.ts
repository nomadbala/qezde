import { z } from "zod";

export const ApiResponseSchema = <T extends z.ZodType>(dataSchema: T) =>
  z.object({
    data: dataSchema,
    status: z.number(),
    message: z.string().optional(),
  });

export type ApiResponse<T> = z.infer<
  ReturnType<typeof ApiResponseSchema<z.ZodType<T>>>
>;

export const ApiErrorResponseSchema = z.object({
  message: z.string(),
  status: z.number(),
  code: z.string().optional(),
});

export type ApiErrorResponse = z.infer<typeof ApiErrorResponseSchema>;
