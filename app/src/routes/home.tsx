import { lazy } from "solid-js";
import { Route } from "@solidjs/router";

const HomeRoute: typeof Route = props => {
  return <Route {...props} component={lazy(() => import("@/pages/home/Home.tsx"))} />;
};

export default HomeRoute;
