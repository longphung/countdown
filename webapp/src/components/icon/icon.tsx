import React, { ReactNode } from "react";
import { Icons } from "./constants";

type Props = {
  icon: Icons;
  className?: string;
  children?: ReactNode;
};

const Icon: React.FC<Props> = (props) => {
  const { icon, className = "", children = null } = props;
  // Render
  return <i className={`icon-${icon} ${className}`}>{children}</i>;
};

export default Icon;
