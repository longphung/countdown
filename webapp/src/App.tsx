import { useState } from "react";
import "./index.css";
import ItemsList from "@components/items_list";
import InputWithError from "@pages/input_with_error";
import Button from "@pages/button";

function App() {
  const [count, setCount] = useState("fsd");
  console.log('hf')
  /* =========
     Render
     ========= */
  return (
    <div className="container mx-auto">
      <InputWithError label={"Count"} value={count} onChange={(e) => setCount(e.target.value)} />
      <Button>hihi</Button>
      <ItemsList />
    </div>
  );
}

export default App;
