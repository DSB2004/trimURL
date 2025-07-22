import Form from "../components/form";
import Display from "../components/display";
import { URLProvider } from "../provider/url.provider";
export default function Shortener() {
  return (
    <URLProvider>
      <div className="bg-white shadow-sm w-full md:w-4/5 mx-auto p-4 flex flex-col gap-5 items-center md:items-start">
        <h2 className="text-xl font-medium">Paste the URL to be shortened</h2>
        <Form></Form>
        <div className="text-center font-normal text-sm md:text-base md:font-medium mx-auto">
          <p>
            URL Shortener is a free tool to shorten URLs and generate short
            links URL
          </p>
          <p>Allows to create a shortened link making it easy to share</p>
        </div>
        <Display></Display>
      </div>
    </URLProvider>
  );
}
