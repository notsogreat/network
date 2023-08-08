import React, { useState } from "react";
import axios from "axios";

const Dashboard: React.FC = () => {
  const [scriptOutput, setScriptOutput] = useState<string | null>(null);
  const handleRunScript = async () => {
    try {
      const response = await axios.get("http://localhost:8080/api/info");
      if (response.status === 200) {
        console.log("Script executed successfully.");
        const output = response.data.output;
        setScriptOutput(output);
      } else {
        console.error("An error occurred while running the script.");
      }
    } catch (error) {
      console.error("An error occurred while calling the API:", error);
    }
  };

  return (
    <div>
      <h1>Dashboard</h1>
      <button onClick={handleRunScript}>Run Bash Script</button>
      {scriptOutput && (
        <div>
          <h2>Script Output:</h2>
          <pre>{scriptOutput}</pre>
        </div>
      )}
    </div>
  );
};

export default Dashboard;
