import type { Component } from "solid-js";
import { Button } from "@/components/ui/button";
import { A } from "@solidjs/router";

const Tracks: Component = () => {
  return (
    <div>
      <h1 class="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl text-center">
        Tracks
      </h1>
      <div class="flex flex-col items-center justify-center h-screen space-y-4">
        <Button variant="link" as={A} href="1">
          Track 1
        </Button>
        <Button variant="link" as={A} href="2">
          Track 2
        </Button>
        <Button variant="link" as={A} href="/">
          Home
        </Button>
      </div>
    </div>
  );
};

export default Tracks;
