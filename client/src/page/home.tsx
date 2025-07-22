import Shortener from "../feature/shortener";
export default function Home() {
  return (
    <div className="min-h-screen bg-gray-200 ">
      <h2 className="text-center p-4 text-4xl font-bold">URL Shortener</h2>
      <Shortener></Shortener>
    </div>
  );
}
