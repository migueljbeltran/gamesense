export type MatchStatus = "uploaded" | "processing" | "completed" | "failed";

export type Match = {
  id: string;
  status: MatchStatus;
  createdAt: string;
};
