import React, { ReactNode } from "react";
import { AnimatePresence, motion } from "framer-motion";
// Local modules
import styles from "./button.module.css";
import { createRipple } from "./utils";

type Props<T> = {
  children?: ReactNode;
  type?: "button" | "submit" | "reset";
  className?: string;
  loading?: boolean;
  // Disable ripples
  noRipple?: boolean;
  // Custom tags
  as?: T;
  // Events
  onClick?: (event: React.MouseEvent<any>) => any;
  onMouseDown?: (event: React.MouseEvent<any>) => any;
  [key: string]: any;
};

const ButtonInner = <T extends string,>(props: Props<T>, ref: React.ForwardedRef<T>) => {
  const {
    children,
    type = "button",
    className = "",
    loading,
    as = "button",
    onMouseDown,
    noRipple = false,
    ...rest
  } = props;
  // Functions
  const handleMouseDown = (event: React.MouseEvent<HTMLElement>): void => {
    onMouseDown && onMouseDown(event);
    !noRipple && createRipple(event);
  };
  /* =========
       Render
       ========= */
  return React.createElement(
    as,
    {
      className: `btn relative overflow-hidden ${styles.button} ${className}`,
      type,
      onMouseDown: handleMouseDown,
      ref,
      ...rest,
    },
    <>
      <AnimatePresence>
        {loading && (
          <motion.span
            key="loader"
            className="overflow-hidden"
            initial={{ width: 0 }}
            animate={{ width: "fit-content" }}
            exit={{ width: 0 }}
          >
            <div className="inline-flex bg-pink-600 text-white rounded-full h-6 px-3 justify-center items-center loading bg-transparent border-none p-0" />
          </motion.span>
        )}
      </AnimatePresence>
      {children}
    </>
  );
};

const Button = React.forwardRef(ButtonInner);

export default Button;
