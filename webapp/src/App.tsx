import { useState } from "react";
import "./index.css";
import Items_list from "@components/items_list";

function App() {
  const [count, setCount] = useState(0);
  console.log('hf')
  /* =========
     Render
     ========= */
  return (
    <div className="App">
      <Items_list />
    </div>
  );
}

export default App;
