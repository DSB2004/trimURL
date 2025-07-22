export const CopyToClipboard = (value: string) => {
  navigator.clipboard
    .writeText(value)
    .then(() => {
      alert("Copied to clipboard successfully");
    })
    .catch((err) => {
      console.error("Failed to copy:", err);
      alert("Failed to copy to clipboard");
    });
};
