import { type Component } from "solid-js";
import { Router } from "@solidjs/router";

import HomeRoute from "@/routes/home";
import TracksRoute from "@/routes/tracks";

const AppRouter: Component = () => {
  return (
    <Router>
      <HomeRoute path="/" />
      <TracksRoute path="/tracks" />
    </Router>
  );
};

export default AppRouter;
