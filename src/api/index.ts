import { Router } from "express";

import users from "./Users/users.route";
const router: Router = Router();

router.use("/users", users);
// router.use("/projects", projects);

export default router;
