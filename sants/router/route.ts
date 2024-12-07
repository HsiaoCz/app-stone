import express from "express";
import handleCreateUser from "../app/user";

const app = express();

app.get("/", handleCreateUser);
app.listen(3000);
