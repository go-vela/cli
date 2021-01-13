// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package login

// authSuccess provides the HTML for rendering
// a message in the browser after successfully
// completing the oauth workflow.
//
// nolint:lll // lots of html, ignoring line length
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
  fill: hsl(286, 29%, 51%);
}
.vela-logo-outer {
  fill: hsl(192, 100%, 50%);
}
.vela-logo-inner {
  fill: hsl(0, 0%, 98%);
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
  .vela-logo-inner {
    fill: hsl(0, 0%, 16%);
  }
  .box {
    background-color: hsl(0, 0%, 98%);
}
</style>
<body>
  <svg width="52" height="52" viewBox="0 0 1500 1500" class="vela-logo"><path class="vela-logo-star" d="M1477.22 329.54l-139.11-109.63 11.45-176.75-147.26 98.42-164.57-65.51 48.11 170.47-113.16 136.27 176.99 6.93 94.63 149.72 61.28-166.19 171.64-43.73z"></path><path class="vela-logo-outer" d="M1174.75 635.12l-417.18 722.57a3.47 3.47 0 01-6 0L125.38 273.13a3.48 3.48 0 013-5.22h796.86l39.14-47.13-14.19-50.28h-821.8A100.9 100.9 0 0041 321.84L667.19 1406.4a100.88 100.88 0 00174.74 0l391.61-678.27z"></path><path class="vela-logo-inner" d="M1087.64 497.29l-49.37-1.93-283.71 491.39L395.9 365.54H288.13l466.43 807.88 363.02-628.76-29.94-47.37z"></path></svg>
  <div class="box">
    <h1>Successfully authenticated with Vela!</h1>
    <p>You may now close this tab and return to the terminal.</p>
  </div>
</body>
`
