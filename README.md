\## Authentication \& Security Checkpoint Handling



This project includes an authentication module designed for \*\*educational and evaluation purposes only\*\*.



During testing, LinkedIn presented a \*\*security verification (reCAPTCHA)\*\* after login submission.  

The automation \*\*detected this checkpoint and halted execution immediately\*\* without attempting any bypass.



\### Why this is intentional

\- Automating CAPTCHA or 2FA violates LinkedIn’s Terms of Service

\- The assignment explicitly requires \*\*detection and graceful handling\*\* of security checkpoints

\- The goal is to demonstrate \*\*engineering judgment\*\*, not circumvention



\### Behavior Implemented

\- Detects invalid credentials

\- Detects 2FA (PIN input)

\- Detects CAPTCHA challenges

\- Stops automation safely and keeps the browser open for inspection



This approach ensures compliance, safety, and correct demonstration of automation best practices.



