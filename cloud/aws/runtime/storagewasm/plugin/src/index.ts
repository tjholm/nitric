export function write() {
  const input = Host.inputString();
  const request = JSON.parse(input);

  console.log(request);

  Host.outputString(JSON.stringify({}));
}
