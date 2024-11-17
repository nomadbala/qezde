import { FC } from "react";
import { IconProps } from "..";
import { cn } from "@/lib/utils";

export const LinkedinIcon: FC<IconProps> = ({
  size = 24,
  height,
  width,
  ...props
}) => {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      fill="none"
      viewBox="0 0 48 48"
      height={height || size}
      width={width || size}
      {...props}
      className={cn(props.className, "rounded-sm")}
    >
      <g clipPath="url(#a)">
        <path
          fill="white"
          d="M44.46 0H3.54A3.54 3.54 0 0 0 0 3.54v40.92A3.54 3.54 0 0 0 3.54 48h40.92A3.54 3.54 0 0 0 48 44.46V3.54A3.54 3.54 0 0 0 44.46 0M14.3 40.89H7.09V17.97h7.22zm-3.62-26.1a4.14 4.14 0 1 1 3.87-2.54 4.1 4.1 0 0 1-3.87 2.54M40.9 40.91h-7.22V28.39c0-3.7-1.57-4.84-3.6-4.84-2.13 0-4.23 1.62-4.23 4.93v12.43h-7.22V17.98h6.94v3.18h.1c.69-1.41 3.13-3.82 6.85-3.82 4.03 0 8.38 2.39 8.38 9.39z"
        />
      </g>
      <defs>
        <clipPath id="a">
          <path fill="#fff" d="M0 0h48v48H0z" />
        </clipPath>
      </defs>
    </svg>
  );
};
