"use client"

import { oswald } from "@/lib/fonts";
import { cn } from "@/lib/utils";
import Link from "next/link";

import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { z } from "zod"

import { toast } from "@/hooks/use-toast";
import { Button } from "@/components/ui/button"
import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"

const FormSchema = z.object({
    email: z
        .string()
        .min(1, { message: "Email has to be filled" })
        .email("Email is not valid"),
    password: z
        .string()
        .min(6, { message: "Password has to be filled and include minimum 6 characters" })
})

export default function LoginPage() {
    const form = useForm<z.infer<typeof FormSchema>>({
        resolver: zodResolver(FormSchema),
    })

    function onSubmit(data: z.infer<typeof FormSchema>) {
        toast({
            title: "You submitted the following values:",
            description: (
                <pre className="mt-2 w-[340px] rounded-md bg-slate-950 p-4">
                    <code className="text-white">{JSON.stringify(data, null, 2)}</code>
                </pre>
            )
        })
    }

    return (
        <div className="w-full h-screen flex justify-center">
            <div className="h-screen w-5/12 bg-[url('/1.svg')] bg-no-repeat bg-center bg-cover flex-col items-center justify-center gap-6 hidden md:flex">
                <svg width="64" height="64" viewBox="0 0 179 185" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <path fill-rule="evenodd" clip-rule="evenodd" d="M135.424 183.04L96.512 149.248C89.344 150.443 82.6027 151.04 76.288 151.04C50.8587 151.04 31.744 145.067 18.944 133.12C6.31467 121.003 0 101.632 0 75.008C0 48.896 6.99733 29.8667 20.992 17.92C34.9867 5.97333 53.504 0 76.544 0C125.013 0 149.248 24.1493 149.248 72.448C149.248 98.9013 142.592 118.869 129.28 132.352L157.696 152.064L135.424 183.04ZM75.264 120.832C80.0427 120.832 84.0533 120.064 87.296 118.528C90.5387 116.821 93.0987 113.92 94.976 109.824C98.56 102.827 100.352 91.392 100.352 75.52C100.352 59.9893 98.56 48.7253 94.976 41.728C91.392 34.56 84.8213 30.8053 75.264 30.464C71.68 30.464 67.9253 30.8907 64 31.744C60.2453 32.4267 57.856 33.4507 56.832 34.816C54.784 37.2053 53.1627 42.3253 51.968 50.176C50.944 57.856 50.432 66.304 50.432 75.52C50.432 89.856 51.7973 100.352 54.528 107.008C56.4053 111.957 59.0507 115.541 62.464 117.76C65.8773 119.808 70.144 120.832 75.264 120.832ZM76.352 83.848C73.152 83.8053 70.3147 83.208 67.84 82.056C67.84 80.136 67.7333 77.9813 67.52 75.592C67.3493 73.16 67.0293 70.8773 66.56 68.744L67.072 68.488C69.4187 68.9573 71.616 69.192 73.664 69.192C75.328 69.192 76.6507 69.0427 77.632 68.744C78.6133 68.4027 79.104 67.912 79.104 67.272C79.104 66.632 78.528 66.1627 77.376 65.864C76.2667 65.5653 74.7947 65.416 72.96 65.416C71.8933 65.416 70.336 65.5013 68.288 65.672L66.88 57.864C67.8613 56.9253 68.8213 56.3067 69.76 56.008C70.7413 55.6667 71.7867 55.496 72.896 55.496C75.4987 55.496 77.7173 56.0933 79.552 57.288C81.4293 58.4827 82.8373 60.04 83.776 61.96C84.7147 63.88 85.184 65.8853 85.184 67.976C85.184 71.2187 84.4373 73.7787 82.944 75.656C81.4933 77.4907 79.3813 78.6 76.608 78.984L76.672 83.528L76.352 83.848ZM72.448 97.096C70.9547 97.096 69.76 96.648 68.864 95.752C67.968 94.8133 67.52 93.5547 67.52 91.976C67.52 90.6107 67.968 89.4587 68.864 88.52C69.8027 87.5813 70.9547 87.112 72.32 87.112C73.728 87.112 74.9013 87.5813 75.84 88.52C76.8213 89.4587 77.312 90.6107 77.312 91.976C77.312 93.5547 76.8427 94.8133 75.904 95.752C75.008 96.648 73.856 97.096 72.448 97.096ZM168.32 184.712C173.843 184.712 178.32 180.235 178.32 174.712C178.32 169.189 173.843 164.712 168.32 164.712C162.797 164.712 158.32 169.189 158.32 174.712C158.32 180.235 162.797 184.712 168.32 184.712Z" fill="#1E1E1E" />
                </svg>
                <h2 className={cn("text-white text-4xl font-semibold", oswald.className)}>С возвращением!</h2>
                <span className={cn("text-white text-center text-xl w-3/5  font-light", oswald.className)}>Введите свои данные и продолжите ваше приключение с нами!</span>
            </div>
            <div className="flex flex-col items-center justify-center gap-5 w-8/12">
                <h2 className={cn("text-accent text-4xl font-semibold", oswald.className)}>Вход в аккаунт</h2>
                <menu className="flex items-center gap-4">
                    <li className="w-9 h-9 flex items-center justify-center rounded-full border-2 cursor-pointer hover:brightness-50 transition-all">
                        <Link href="#">
                            <svg width="16" height="16" viewBox="0 0 7 15" fill="none" xmlns="http://www.w3.org/2000/svg">
                                <path d="M6.53532 5.13372H4.31765V3.70694C4.31765 3.17111 4.67967 3.04619 4.93465 3.04619C5.18905 3.04619 6.49964 3.04619 6.49964 3.04619V0.69061L4.34433 0.682358C1.95174 0.682358 1.40725 2.43923 1.40725 3.56352V5.13372H0.0235596V7.56101H1.40725C1.40725 10.6761 1.40725 14.4294 1.40725 14.4294H4.31765C4.31765 14.4294 4.31765 10.6391 4.31765 7.56101H6.2815L6.53532 5.13372Z" fill="#565656" />
                            </svg>
                        </Link>
                    </li>
                    <li className="w-9 h-9 flex items-center justify-center rounded-full border-2 cursor-pointer hover:brightness-50 transition-all">
                        <Link href="#">
                            <svg width="16" height="16" viewBox="0 0 17 17" fill="none" xmlns="http://www.w3.org/2000/svg">
                                <path d="M16.3054 7.42646C13.7578 7.42524 11.2102 7.42585 8.66262 7.42614C8.66295 8.46162 8.66136 9.49711 8.66324 10.5323C10.1386 10.5319 11.614 10.5316 13.0891 10.5323C12.9181 11.5245 12.3164 12.4319 11.463 12.9906C10.9266 13.344 10.3103 13.5734 9.67485 13.6825C9.03535 13.7895 8.37456 13.8031 7.73664 13.6766C7.08794 13.5497 6.46871 13.2849 5.92851 12.9116C5.06473 12.3169 4.40552 11.4468 4.06706 10.4676C3.72096 9.47158 3.71842 8.36334 4.06864 7.3683C4.31164 6.66699 4.71609 6.01916 5.24741 5.49195C5.90279 4.83448 6.75485 4.36447 7.67287 4.17173C8.45894 4.00729 9.28782 4.03867 10.0574 4.26684C10.7115 4.46145 11.3146 4.81271 11.8085 5.27495C12.3078 4.78844 12.8045 4.29915 13.3029 3.81174C13.5643 3.54938 13.839 3.29851 14.0921 3.02901C13.3353 2.34261 12.448 1.79085 11.4786 1.44206C9.73321 0.812598 7.76487 0.799238 6.0078 1.39518C4.02771 2.05949 2.34168 3.51056 1.40686 5.34466C1.0814 5.97665 0.843782 6.6512 0.701041 7.34504C0.341942 9.07371 0.592235 10.9168 1.4059 12.4913C1.93471 13.519 2.69285 14.4326 3.61249 15.1489C4.4801 15.8268 5.49137 16.3277 6.56358 16.6071C7.91652 16.9627 9.35668 16.9547 10.7185 16.651C11.9493 16.3734 13.1139 15.797 14.0436 14.9558C15.0264 14.0708 15.7274 12.9048 16.0986 11.6517C16.5033 10.2849 16.5592 8.82565 16.3054 7.42646Z" fill="#565656" />
                            </svg>
                        </Link>
                    </li>
                    <li className="w-9 h-9 flex items-center justify-center rounded-full border-2 cursor-pointer hover:brightness-50 transition-all">
                        <Link href="#">
                            <svg width="16" height="16" viewBox="0 0 17 16" fill="none" xmlns="http://www.w3.org/2000/svg">
                                <path d="M1.0489 5.17727H4.46024V15.4294H1.0489V5.17727ZM2.7774 0.235291C1.60984 0.235291 0.847046 1.00078 0.847046 2.00551C0.847046 2.98958 1.58752 3.77732 2.73276 3.77732H2.75451C3.94439 3.77732 4.6854 2.98954 4.6854 2.00551C4.66307 1.00078 3.94443 0.235291 2.7774 0.235291ZM12.836 4.93644C11.0252 4.93644 10.214 5.9311 9.76146 6.62865V5.17727H6.34909C6.39427 6.13904 6.34909 15.4294 6.34909 15.4294H9.76146V9.70389C9.76146 9.39724 9.78378 9.09173 9.87356 8.87205C10.1206 8.25989 10.6821 7.62596 11.6238 7.62596C12.8594 7.62596 13.3529 8.56654 13.3529 9.94412V15.4294H16.7647V9.5506C16.7647 6.40163 15.0819 4.93644 12.836 4.93644Z" fill="#565656" />
                            </svg>
                        </Link>
                    </li>
                </menu>
                <span className={cn("font-light text-muted", oswald.className)}>или используйте <u>почту</u> для входа</span>
                <Form {...form}>
                    <form onSubmit={form.handleSubmit(onSubmit)} className="w-5/6 md:w-2/3 space-y-6 flex-col">
                        <FormField
                            control={form.control}
                            name="email"
                            render={({ field }) => (
                                <FormItem>
                                    <FormLabel>Почта</FormLabel>
                                    <FormControl>
                                        <Input {...field} type="email"/>
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                        <FormField
                            control={form.control}
                            name="password"
                            render={({ field }) => (
                                <FormItem>
                                    <FormLabel>Пароль</FormLabel>
                                    <FormControl>
                                        <Input {...field} type="password"/>
                                    </FormControl>
                                    <FormMessage />
                                </FormItem>
                            )}
                        />
                        <Button type="submit" className="bg-accent hover:bg-accent hover:brightness-90 transition-all">Войти</Button>
                    </form>
                </Form>
            </div>
        </div>
    )
}