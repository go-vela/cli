// SPDX-License-Identifier: Apache-2.0

package repo

// authSuccess provides the HTML for rendering
// a message in the browser after successfully
// completing the oauth workflow.
const authSuccess = `
<!doctype html>
<meta charset="utf-8">
<title>Success: Vela CLI</title>
<style type="text/css">
  body {
    color: hsl(0, 0%, 98%);
    background-color: hsl(0, 0%, 16%);
    font-size: 14px;
    font-family: -apple-system, "Segoe UI", Helvetica, Arial, sans-serif;
    line-height: 1.5;
    max-width: 620px;
    margin: 28px auto;
    text-align: center;
  }

  .vela-logo-star {
    fill: hsl(289, 54.8%, 57.5%);
  }

  .vela-logo-lines {
    fill: hsl(194, 89.7%, 58%);
  }

  .box {
    background-color: hsl(0, 0%, 16%);
  }

  h1 {
    font-size: 24px;
    margin-bottom: 0;
  }

  p {
    margin-top: 0;
  }

  .box {
    border: 1px solid hsl(286, 29%, 51%);
    padding: 24px;
    margin: 28px;
  }

  @media (prefers-color-scheme: light) {
    body {
      color: hsl(0, 0%, 16%);
      background-color: hsl(0, 0%, 98%);
    }

    .box {
      background-color: hsl(0, 0%, 98%);
    }
</style>

<body>
  <svg width="72" height="72" viewBox="0 0 1920 1920" class="vela-logo">
    <path class="vela-logo-lines"
      d="M618.73 431.74h-162.1a56.87 56.87 0 0 0-50.86 82.3l501.05 1002.1a56.86 56.86 0 0 0 101.72 0l332.85-665.72 63.63 127.07-294.74 589.5a170.64 170.64 0 0 1-152.6 94.33 170.64 170.64 0 0 1-152.6-94.33L304.03 564.9A170.61 170.61 0 0 1 456.63 318h105.14l56.96 113.75Z" />
    <path class="vela-logo-lines"
      d="M625.05 318h126.9l56.94 113.74h-126.9L625.05 318Zm253.65 0h63.45l56.94 113.74h-63.44L878.7 318ZM675.82 545.47l281.86 563.74 147.3-294.62 137.58-20.82-284.88 569.76-409.03-818.06h127.17Z" />
    <path class="vela-logo-star"
      d="m1372.75 659.05-234.4 35.44 168.8-166.43L1201.96 318l209.51 107.16 168.8-166.45-38.7 233.88 210.46 109.1-234.4 35.43-38.7 233.89-106.17-211.97Z" />
  </svg>
  <div class="box">
    <h1>Successfully authenticated with Vela!</h1>
    <p>You may now close this tab and return to the terminal.</p>
  </div>
</body>
`
