import { useState } from "react";
import { Greet } from "../wailsjs/go/desktop/App";
import { Sidebar } from "flowbite-react";
import {
  RectangleStackIcon,
  MagnifyingGlassIcon,
} from "@heroicons/react/24/solid";

function App() {
  const [resultText, setResultText] = useState(
    "Please enter your name below aha👇"
  );
  const [name, setName] = useState("");
  const updateName = (e: any) => setName(e.target.value);
  const updateResultText = (result: string) => setResultText(result);

  function greet() {
    Greet(name).then(updateResultText);
  }

  return (
    <div className="flex">
      <div className="flex-none dark w-fit h-screen">
        <Sidebar aria-label="App Sidebar">
          <Sidebar.Items>
            <Sidebar.ItemGroup>
              <Sidebar.Item href="#" icon={MagnifyingGlassIcon}>
                Lens
              </Sidebar.Item>
              <Sidebar.Item
                href="#"
                icon={RectangleStackIcon}
                label="5"
                labelColor="alternative"
              >
                Tasks
              </Sidebar.Item>
            </Sidebar.ItemGroup>
          </Sidebar.Items>
        </Sidebar>
      </div>

      <div className="flex-auto">
        <div className="flex flex-col h-screen">
          <div className="flex-none h-1/2">Deletion List</div>
          <div className="flex-none h-1/2 bg-green-300">File List</div>
        </div>
      </div>
    </div>
  );
}

export default App;
