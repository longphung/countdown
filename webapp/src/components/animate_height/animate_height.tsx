import React from "react";
import { AnimatePresence, motion } from "framer-motion";

type Props = {
  children: React.ReactNode;
  isVisible: boolean;
  className?: string;
  collapsed?: {
    marginBottom: number | string;
  };
  visible?: {
    marginBottom: number | string;
  };
};

const AnimateHeight: React.FC<Props> = (props) => {
  const {
    children,
    isVisible,
    className,
    collapsed = { marginBottom: 0 },
    visible = { marginBottom: "1.5rem" },
  } = props;
  /* =========
     Render
     ========= */
  return (
    <AnimatePresence>
      {isVisible && (
        <motion.div
          className={className}
          initial={{ opacity: 0, height: 0, x: 30, ...collapsed }}
          animate={{ height: "auto", opacity: 1, x: 0, ...visible }}
          exit={{
            height: 0,
            opacity: 0,
            x: 30,
            ...collapsed,
          }}
        >
          {children}
        </motion.div>
      )}
    </AnimatePresence>
  );
};

export default AnimateHeight;
