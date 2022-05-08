import React from "react";
import { formatDistanceToNow } from "date-fns";
// Local modules
import { Countdown } from "@services/types";

type Props = {
  item: Pick<Countdown, "name" | "dueDate">;
};

const Item: React.FC<Props> = (props) => {
  const { item } = props;
  /* =========
     Render
     ========= */
  return (
    <tr className="hover:bg-gray-100 dark:hover:bg-gray-700">
      <td className="py-4 px-6 text-sm font-medium text-gray-900 whitespace-nowrap dark:text-white">{item.name}</td>
      <td className="capitalize py-4 px-6 text-sm font-medium text-gray-500 whitespace-nowrap dark:text-white">
        {formatDistanceToNow(new Date(item.dueDate))}
      </td>
      <td className="py-4 px-6 text-sm font-medium text-right whitespace-nowrap">
        <a href="#" className="text-blue-600 dark:text-blue-500 hover:underline">
          Edit
        </a>
      </td>
    </tr>
  );
};

export default Item;
