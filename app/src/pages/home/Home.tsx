import type { Component } from "solid-js";
import { Button } from "@/components/ui/button";
import { A } from "@solidjs/router";

const Home: Component = () => {
  return (
    <div>
      <h1 class="scroll-m-20 text-4xl font-extrabold tracking-tight lg:text-5xl text-center">
        Home
      </h1>
      <div class="flex flex-col items-center justify-center h-screen space-y-4">
        <Button variant="link" as={A} href="/tracks">
          Tracks
        </Button>
        <Button variant="link" as={A} href="/settings">
          Settings
        </Button>
      </div>
    </div>
  );
};

export default Home;
