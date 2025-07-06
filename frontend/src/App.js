import React, { useState, useEffect } from "react";
import CodeMirror from "@uiw/react-codemirror";
import { vscodeDark } from "@uiw/codemirror-theme-vscode";
import { json } from "@codemirror/lang-json";
import ReactJson from "react-json-view";

function App() {
  const [requestJson, setRequestJson] = useState('');
  const [mappingJson, setMappingJson] = useState('');
  const [output, setOutput] = useState(null);

  // 🔁 Reusable mapping function
  const mapJson = (jsonObj, mapping) => {
    const result = {};
    for (let key in mapping) {
      const path = mapping[key];
      if (!path) {
        result[key] = "";
        continue;
      }
      const parts = path.split(".");
      let value = jsonObj;
      for (let part of parts) {
        if (value && part in value) {
          value = value[part];
        } else {
          value = null;
          break;
        }
      }
      result[key] = value;
    }
    return result;
  };

  useEffect(() => {
  try {
    const parsedJson = JSON.parse(requestJson);
    const parsedMapping = JSON.parse(mappingJson);
    const result = mapJson(parsedJson, parsedMapping);
    setOutput(result);
  } catch (err) {
    setOutput(null); // Hide output if there's an error
  }
}, [requestJson, mappingJson]); // Trigger whenever either changes


  // 🌐 Send to backend (optional button)
  const sendToBackend = async () => {
    try {
      const res = await fetch("http://localhost:8080/api/map-json", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          requestJson: JSON.parse(requestJson),
          requestMapping: JSON.parse(mappingJson),
        }),
      });
        const data = await res.json();
        setOutput(data);
      
    } catch (err) {
      alert("Error: " + err.message);
    }
  };


  const handleRequestJsonUpload = (e) => {
  const file = e.target.files[0];
  if (!file) return;
  const reader = new FileReader();
  reader.onload = (event) => {
    try {
      const jsonText = event.target.result;
      JSON.parse(jsonText); // Check if valid
      setRequestJson(jsonText);
    } catch (err) {
      alert("Invalid JSON in uploaded file.");
    }
  };
  reader.readAsText(file);
};

const handleMappingJsonUpload = (e) => {
  const file = e.target.files[0];
  if (!file) return;
  const reader = new FileReader();
  reader.onload = (event) => {
    try {
      const jsonText = event.target.result;
      JSON.parse(jsonText); // Check if valid
      setMappingJson(jsonText);
    } catch (err) {
      alert("Invalid Mapping JSON in uploaded file.");
    }
  };
  reader.readAsText(file);
};



  return (
    <div style={{ padding: 20 }}>
      <h2>🛠 Dynamic JSON Mapper</h2>

      <div style={{ display: "flex", gap: "20px", marginBottom: "20px" }}>
        <div style={{ flex: 1 }}>
          <h4>📥 Request JSON</h4>
          <CodeMirror
            value={requestJson}
            height="250px"
            extensions={[json()]}
            theme={vscodeDark}
            onChange={(val) => {
              setRequestJson(val);
              try {
                const parsedJson = JSON.parse(val);
                const parsedMapping = JSON.parse(mappingJson);
                const result = mapJson(parsedJson, parsedMapping);
                setOutput(result);
              } catch (err) {
                setOutput(null); // Clear output if JSON invalid
              }
            }}
          />
        </div>

        <input type="file" accept=".json" onChange={handleRequestJsonUpload} />

        <div style={{ flex: 1 }}>
          <h4>🧭 Request Mapping</h4>
          <CodeMirror
            value={mappingJson}
            height="250px"
            extensions={[json()]}
            theme={vscodeDark}
            onChange={(val) => {
              setMappingJson(val);
              try {
                const parsedJson = JSON.parse(requestJson);
                const parsedMapping = JSON.parse(val);
                const result = mapJson(parsedJson, parsedMapping);
                setOutput(result);
              } catch (err) {
                setOutput(null); // Clear output if mapping invalid
              }
            }}
          />
        </div>

        <input type="file" accept=".json" onChange={handleMappingJsonUpload} />
      </div>

      <div style={{ display: "flex", gap: "10px", marginBottom: "20px" }}>
        <button onClick={sendToBackend}>Map via Backend API</button>
      </div>

      <div>
        <h4>🔍 Mapped Output:</h4>
        <ReactJson
          src={output || {}}
          collapsed={false}
          displayDataTypes={false}
          name={false}
          style={{padding: "10px", borderRadius: "8px" }}
        />
      </div>
    </div>
  );
}

export default App;
