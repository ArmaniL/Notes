import * as React from "react";
import LoginForm from "./components/LoginForm";
import {
    BrowserRouter as Router,
    Route,
    Routes,

} from "react-router-dom";
import HomePage from "./components/HomePage";
import NewNoteForm from "./components/NewNoteForm";

export default function App() {
    const [token, setToken] = React.useState<string>("")
    return <Router>
        <Routes>
            <Route path="/" element={<LoginForm authControls={setToken}></LoginForm>}></Route>
            <Route path="/dashboard" element={<HomePage token={token} />}></Route>
            <Route path="/new-note" element={<NewNoteForm token={token} />}></Route>

        </Routes>
    </Router>
}