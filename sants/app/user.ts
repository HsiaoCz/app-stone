import { Request, Response } from "express";

const handleCreateUser = (req: Request, res: Response) => {
  res.send("hello");
};

export default handleCreateUser;
