# The Clean Coder: A Code of Conduct for Professional Programmers

## 8. Testing Strategies

- typically not executed as part of the Continuous Integration suite, because they often have longer runtimes
- at this level that we might see performance and throughput tests.
- Integration tests are choreography tests. They do not test business rules
- The test automation pyramid
- In general, the business writes the happy-path tests, while QA writes the corner, boundary, and unhappy-path tests.
- it should be the goal of the development group that QA find nothing wrong.

## 2. Saying No

- “Is good code impossible? Is professionalism impossible?”

Answer: I say no
- THE COST OF SAYING YES
- he is really asking “Is
- said one thing correctly all along: The code was a throwaway
- The client will always extend the deadline. They will always want more features.
- Good code has become impossible.
- Is good code impossible in modern software development?
- You are lying
- trying
- Perhaps you don’t like that idea? Perhaps you think trying is a positive thing to do. After all, would Columbus have discovered America if he hadn’t tried?
- The higher the stakes, the more valuable no becomes
- Providing too much detail can be an invitation for micro-management.
- can
- Managers are people with a job to do, and most managers know how to do that job pretty well. Part of that job is to pursue and defend their objectives as aggressively as they can.

By the same token, programmers are also people with a job to do, and most of them know how to get that job done pretty well. If they are professionals they will pursue and defend their objectives as aggressively as they can.
- Indeed
- No. Not if you are a professional.

Slaves are not allowed to say no. Laborers may be hesitant to say no. But professionals are expected to say no

## 1. Professionalism

- Laugh
- He will not demean another for making a mistake, because he knows he may be the next to fail
- timid
- programming is an act of supreme arrogance.
- put yourself in your employer’s shoes 
- You don’t have to be a domain expert, but there is a reasonable amount 
- Mentoring
- simple exercises such as the Bowling Game or Prime Factors. I call these exercises kata. There are many such kata
- Here is a minimal list of the things that every software professional should be conversant with:
- You should plan on working 60 hours per week. The first 40 are for your employer. The remaining 20 are for you. During this remaining 20 hours you should be reading, practicing, learning, and otherwise enhancing your career.
- you refine the design so that the next change is easier
- surprised

## 10. Estimation

- Estimates are fraught with error. That’s why they are called estimates.
- We don’t want people changing their estimates based on what they see other people do.
- estimation technique called “wideband delphi.
- They can help you estimate your tasks more accurately than you can estimate them on your own
- Software professionals are very careful to set reasonable expectations despite the pressure to try to go fast.
- standard deviation4
- When you estimate a task, you provide three numbers
- PERT is the way that estimates are calculate
- try is a loaded term
- An estimate is not a number. An estimate is a distribution.
- Probability distribution
- An estimate is a guess. No commitment is implied
- Missing a commitment is an act of dishonesty only slightly less onerous than an overt lie.
- Professionals don’t make commitments unless they know they can achieve them
- Business likes to view estimates as commitments. Developers like to view estimates as guesses. The difference is profound.
- It took three months
- If you’ve read my books, or heard my talks, you know I rant a lot about independent deployability. This is where I first learned that lesson.

## 11. Pressure

- when you see someone else who’s under pressure, offer to pair with them. Help them out of the hole they are in.
- Pair! When the heat is on, find an associate who is willing to pair program with you
- Rely on Your Disciplines
When the going gets tough, trust your disciplines. The reason you have disciplines is to give you guidance through times of high pressure.
- Avoid creating surprises
- Plot a course to the best possible outcome, and then drive towards that outcome at a reasonable and steady pace
- Manage your stress. Sleepless nights won’t help you get done any faster. Sitting and fretting won’t help either
- if you change your behavior in a crisis, then you don’t truly believe in your normal behavior.
- This does not mean that we spend endless hours polishing code. 
- we don’t tolerate messes. We know that messes will slow us down, causing us to miss dates and break commitments. So we do the best work we can and keep our output as clean as we can.
- In the end, if we can find no way to meet the promises made by the business, then the people who made the promises must accept the responsibility.
- . I laughed at the man in the mirror, the poor schmuck who’d been making life miserable for himself and others in the name of—what?
- In 1988 I was working at Clear Communications. This was a start-up that never quite got started

## 9. Time Management

- They keep their options open by keeping an open mind about alternate solutions. They never become so vested in a solution that they can’t abandon it.
- the inflection point! You can still go back and fix the design. But you can also continue to go forward
- eventually you will make a decision that leads to a mess.
- Professionals avoid getting so vested in an idea that they can’t abandon it and turn around. They keep an open mind about other ideas so that when they hit a dead end they still have other options.
- defense to protect us from the judgment of others.
- preparing for the lie we’ll tell when someone asks us what we are doing and why we are doing it
- priority inversion. You raise the priority of a task so that you can postpone the task that has the true priority.
- 25-minute window of productive time that you aggressively defend against all interruptions.
- Some people get so comfortable with the technique that they estimate their tasks in tomatoes and then measure their weekly tomato velocity
- On a good day you might get 12 or even 14 tomatoes done
- your time is divided into tomato and non-tomato time
- While I ride I listen to podcasts about astronomy or politics. Sometimes I just listen to my favorite music. And sometimes I just turn the headphones off and listen to nature.
- muscle focus helps to recharge mental focu
- you can’t force the focus
- Seven hours of sleep will often give me a full eight hours’ worth of focus-manna.
- Caffeine also puts a strange “jitter” on your focus. 
- Professional developers learn to manage their time to take advantage of their focus-manna.
- Any argument that can’t be settled in five minutes can’t be settled by arguing.” The reason it goes on so long is that there is no clear evidence supporting either side. The argument is probably religious, as opposed to factual.
- Iteration Planning Meetings
- When the meeting gets boring, leave.
- You need to use your time wisely. So be very carefu
- Some in attendance may find them invaluable; others may find them redundant or useless.
- There are two truths about meeting.

Meetings are necessary.
Meetings are huge time wasters.

## 7. Acceptance Testing

- effectively eliminate communication errors between programmers and stakeholders is to write automated acceptance tests
- All too often both parties agree that they understand and leave with completely different ideas
- Every time someone commits a module, the CI system should kick off a build, and then run all the tests in the system. The results of that run should be emailed to everyone on the team.
- Keep the GUI tests to a minimum
- The trick is to design the system so that you can treat the GUI as though it were an API rather than a set of buttons, sliders, grids, and menus
- Unit tests dig into the guts of the system making calls to methods in particular classes. Acceptance tests invoke the system much farther out, at the API or sometimes even UI leve
- Acceptance tests are written by the business for the business (even when you, the developer, end up writing them).
- Acceptance tests are not unit tests. Unit tests are written by programmers for programmers
- Following the principle of “late precision,” acceptance tests should be written as late as possible, typically a few days before the feature is implemented
- done
- The cost of running the manual test plan is so enormous that they have decided to sacrifice it and simply live with the fact that they won’t know if half of their product works
- Acceptance tests should always be automated
- purpose of acceptance tests is communication, clarity, and precisio
- ambiguity of “done.” When a developer says he’s done with a task, what does that mean? Is the developer done in the sense that he’s ready to deploy the feature with full confidence? Or does he mean that he’s ready for QA? Or perhaps he’s done writing it and has gotten it to run once but hasn’t really tested it yet.
- Professional developers don’t flesh out a requirement until they are just about to develop it
- Professional developers understand that estimates can, and should, be made based on low precision requirements, and recognize that those estimates are estimates. To reinforce this, professional developers always include error bars with their estimates so that the business understands the uncertainty.
- In the end, the more precise you make your requirements, the less relevant they become as the system is implemented
- Once they see the requirement actually running, they have a better idea of what they really want—and it’s usually not what they are seeing.
- bit by bit I found myself implementing this application while he sat there and watched. Within the first twenty minutes it was clear that his emphasis had changed from learning how to do it himself to making sure that what I did was what he wanted.
- My explanation simply made no sense to him at all

## 6. Practicing

- Professional programmers often suffer from a lack of diversity in the kinds of problems that they solve. Employers often enforce a single language, platform, and domain in which their programmers must work.
- Practicing a suite of katas is a good way to learn hot keys and navigation idioms. It is also a good way to learn disciplines such as TDD and CI. But most importantly, it is a good way to drive common problem/solution pairs into your subconscious, so that you simply know how to solve them when facing them in real programming.
- hundreds, perhaps thousands, of times. I got very good at it! I could do it in my sleep. I minimized the keystrokes, tuned the variable names, and tweaked the algorithm structure until it was just right. Although I didn’t know it at the time, this was my first kata.
- 60s I could wait a day or two to see the results of a compile. In the late ’70s a 50,000-line program might take 45 minutes to compile. Even in the ’90s, long build times were the norm.

## 5. Test Driven Development

- writing your tests first, creates a force that drives you to a better decoupled design
- There have been several reports4 and studies5 that describe significant defect reduction. From IBM, to Microsoft, from Sabre to Symantec, company after company and team after team have experienced defect reductions of 2X, 5X, and even 10
- You are not allowed to write any production code until you have first written a failing unit test.
You are not allowed to write more of a unit test than is sufficient to fail—and not compiling is failing.
You are not allowed to write more production code that is sufficient to pass the currently failing unit test.

## 4. Coding

- and in general prove to ourselves that we have brains the size of a planet, all while not having to interact with the messy complexities of other people
- No matter how skilled you are, you will certainly benefit from another programmer’s thoughts and ideas.
- There is no way to rush. You can’t make yourself code faster. You can’t make yourself solve problems faster. If you try, you’ll just slow yourself down and make a mess that slows everyone else down, too.
- Your original estimates are more accurate than any changes you make while your boss is confronting you
- You miss elegant solutions because the creative part of your mind is suppressed by the intensity of your focus
- Likewise, a software developer who creates many bugs is acting unprofessionally.
- Test Driven Development (TDD),
- and
- I fixed the problem, of course, but never had the courage to turn off the automatic hack that inspected and fixed the counters. To this day I’m not convinced there wasn’t another hole
- from scratch in assembler
- creative output depends on creative input.
- As soon as you sit down next to someone else, the issues that were blocking you melt away
- that next time you may be the one who needs to interrupt someone else. So the professional attitude is a polite willingness to be helpful.
- Avoid the Zone. This state of consciousness is not really hyper-productive and is certainly not infallible. It’s really just a mild meditative state in which certain rational faculties are diminished in favor of a sense of speed.
- When I am worried about an argument with my wife, or a customer crisis, or a sick child, I can’t maintain focus
- I find myself with my eyes on the screen and my fingers on the keyboard, doing nothing
- working at 3 AM is what serious professionals do. How wrong I was!

That code came back to bite us over and over
- The bottom line is that the opportunity for distraction is high.
- I realized, during that long night, that typing blind is all about confidence
- Being able to sense your errors is really important

## 3. Saying Yes

- Professionals are not required to say yes to everything that is asked of them. However, they should work hard to find creative ways to make “yes” possible
- Years of experience have taught us that breaking disciplines only slows us down.
- Creating a language of commitment may sound a bit scary, but it can help solve many of the communication problems
- telltale signs of noncommitment:

## 12. Collaboration

- the group average still tends in the direction I stated. We, programmers, enjoy the mild sensory deprivation and cocoonlike immersion of focus

