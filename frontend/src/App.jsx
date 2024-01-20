import {useState} from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import {ProcessPullRequest} from "../wailsjs/go/main/App.js";

function App() {
    const [resultText, setResultText] = useState("Please enter a GitHub pull request URL 👇");
    const [name, setName] = useState('');
    const updateName = (e) => setName(e.target.value);
    const updateResultText = (result) => setResultText(result);

    function processPullRequest() {
        ProcessPullRequest(name).then(updateResultText);
    }

    return (
        <div id="App">
            <img src={logo} id="logo" alt="logo"/>
            <div id="result" className="result">{resultText}</div>
            <div id="input" className="input-box">
                <input id="name" className="input" onChange={updateName} autoComplete="off" name="input" type="text"/>
                <button className="btn" onClick={processPullRequest}>Review</button>
            </div>
        </div>
    )
}

export default App
