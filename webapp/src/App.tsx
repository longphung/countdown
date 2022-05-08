import { useState } from "react";
import "./index.css";
import ItemsList from "@components/items_list";
import InputWithError from "@pages/input_with_error";

function App() {
  const [count, setCount] = useState("fsd");
  console.log('hf')
  /* =========
     Render
     ========= */
  return (
    <div className="App">
      <InputWithError label={"Count"} value={count} onChange={(e) => setCount(e.target.value)} />
      <ItemsList />
    </div>
  );
}

export default App;
