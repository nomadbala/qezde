import { FC } from "react";
import { IconProps } from "..";

export const TelegramIcon: FC<IconProps> = ({
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
    >
      <g clipPath="url(#a)">
        <path fill="url(#b)" d="M24 48a24 24 0 1 0 0-48 24 24 0 0 0 0 48" />
        <path
          fill="#fff"
          fillRule="evenodd"
          d="m10.9 23.7 14-6c6.6-2.8 8-3.2 9-3.3q.3 0 .8.3.4.4.4.7v1c-.3 3.7-1.9 13-2.7 17.2q-.6 2.5-1.6 2.4c-1.4.2-2.5-.9-3.8-1.8l-5.4-3.6c-2.4-1.5-.8-2.4.5-3.8.4-.4 6.5-6 6.6-6.4v-.5H28q-.4 0-10 6.7-1.6 1-2.7 1c-.9 0-2.5-.6-3.7-1-1.5-.4-2.7-.7-2.6-1.5q0-.7 1.8-1.4"
          clipRule="evenodd"
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
