import { lazy } from "solid-js";
import { Route } from "@solidjs/router";

const TracksRoute: typeof Route = props => {
  return (
    <Route {...props}>
      <Route path="/" component={lazy(() => import("@/pages/tracks/Tracks.tsx"))} />
      <Route path="/:id" component={lazy(() => import("@/pages/tracks/Track.tsx"))} />
    </Route>
  );
};

export default TracksRoute;
