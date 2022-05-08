import { useQuery } from "react-query";
// Local modules
import { baseInstance } from "../axios_instance";
import { Countdown } from "@services/types";

/* =========
  This module exports:
  NOTE: Queries
  - useGetAllCountdowns
  NOTE: Mutations
========= */

export const useGetAllCountdowns = () => {
  return useQuery<Countdown[], Error>(["countdowns"], async () => {
    const response = await baseInstance.get("/countdowns");
    return response.data;
  });
};
