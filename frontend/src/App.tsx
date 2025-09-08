import { useEffect, useState } from "react";
import "./App.css";
import MyChart from "./components/MyChart";

function App() {
	const [text, setText] = useState("");

	useEffect(() => {
		fetch("/api/ping")
			.then((res) => res.text())
			.then(setText);
	}, []);

	return (
		<>
			<p>{text}</p>
			<MyChart apiUrl="/api/asset/listWithExchangeRate?currency_type=usd" />
		</>
	);
}

export default App;
