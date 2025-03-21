import type { Component } from "solid-js";
import { Button } from "@/components/ui/button";
import { A, useParams } from "@solidjs/router";

const TrackComponent: Component = () => {
  const params = useParams();
  return (
    <div>
      <h1 class="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl text-center">
        Track {params.id}
      </h1>
      <div class="flex flex-col items-center justify-center h-screen space-y-4">
        <Button variant="link" as={A} href="..">
          Back
        </Button>
        <Button variant="link" as={A} href="/">
          Home
        </Button>
      </div>
    </div>
  );
};

export default TrackComponent;
