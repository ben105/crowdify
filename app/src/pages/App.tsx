import { createSignal, Show } from "solid-js";
import { Button } from "@/components/ui/button";
import { isEvening } from "../utils/example";

function App() {
  const [count, setCount] = createSignal(0);

  return (
    <>
      <div class="flex flex-col items-center justify-center h-screen space-y-4">
        <Button onClick={() => setCount(count => count + 1)}>Default - Count: {count()}</Button>
        <Button variant="secondary">Secondary</Button>
        <Button variant="outline">Outline</Button>
        <Button variant="destructive">Destructive</Button>
        <Button variant="ghost">Ghost</Button>
        <Button variant="link">Link</Button>
        <Show when={isEvening(new Date())}>
          <p>Have a good night!</p>
        </Show>
      </div>
    </>
  );
}

export default App;
