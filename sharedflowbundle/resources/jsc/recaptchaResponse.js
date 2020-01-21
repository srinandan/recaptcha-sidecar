 var recaptchaResponse = JSON.parse(context.getVaraible("recaptchaResponse.content"));
 
 //check if riskAnalysis exists
 if (recaptchaResponse.hasOwnProperty("riskAnalysis")) {
     context.setVariable("recaptchaDecision", "true");
     context.setVariable("riskScore", recaptchaResponse.riskAnalysis);
 } else if (recaptchaResponse.hasOwnProperty("tokenProperties")) {
     context.setVariable("recaptchaDecision", "false");
     context.setVariable("riskScore", "0");
     context.setVariable("invalidReason", recaptchaResponse.tokenProperties.invalid_reason);
 } else {
     context.setVariable("recaptchaDecision", "unavailable");
     context.setVariable("riskScore", "0");
 }