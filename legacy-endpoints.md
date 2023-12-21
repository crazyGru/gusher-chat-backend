# Entities


* User
  * Create
  * Follow User
  * Update Profile
  * Activate
  * Ban

* Startup
  * Create
  * Approve
  * Publish
  * Update Avatar
  * Sign ESA
* Applicant
  * Apply
  * Join Startup
  * Sign ESA
  * Rate
* Role
* Goal
* Protocol
  * Assign to startup
* ProtocolTask
  * Assign to Role
* Post
* Channel
* Attachment

* Article
* ArticleTag

* AffiliateAgreement
* AffiliateReferral

* FlowTask

* Topic

* Interest

* Skill

* Notification

# Endpoints

## User

````
/people/follow/:id
/people/following/:id
/people/search
/people/admin/search
/users/search/:keyword/p/:page/:items
/users/search/offset
/users/:id
/users/list/:list/:profile/p/:page/:items
/settings/profile/update
 
````

## Search

````
/search/:filter/:keyword/p/:page/:items/t/:time/:newer
/search/:filter//p/:page/:items/t/:time/:newer
 
````

## Startup

````
/startups/save
/startups/session/save
/startups/avatar-url
/startups/cover-url
/startups/follow/:id
/startups/following/:id
/startups/award/:id/update
/startups/list/:profile/p/:page/:items
/startups/fetch/:status/p/:page/:items
/startups/applicants/:id/p/:page/:items/t/:time/:newer
/startups/search
/startups/admin/search
/roles/fetch/:status/p/:page/:items
/roles/:id/applicants/p/:page/:items/t/:time/:newer
/roles/search
/roles/admin/search
/applications/fetch/:status/p/:page/:items
/applicants/:id/rate/:score
/applicants/:id/scores/p/:page/:items/t/:time/:newer
/startup-internal/:startupId
/startup-internal/:startupId/user-applications/:userId
/agreements/hellosign/callback

````
# Post

````
/posts/new
/posts/all/p/:page/:items/t/:time/:newer
/posts/u/:profile/p/:page/:items/t/:time/:newer
/posts/s/:startup/p/:page/:items/t/:time/:newer
/posts/s/:startup/r/:role/p/:page/:items/t/:time/:newer
/posts/s/:startup/r/:role/a/:applicant/p/:page/:items/t/:time/:newer
/posts/s/:startup/r/:role/g/:goal/p/:page/:items/t/:time/:newer
/messages/:contact/p/:page/:items/t/:time/:newer
/posts/c/:channel/p/:page/:items/t/:time/:newer
/posts/:id/comments/p/:page/:items/t/:time/:newer
/posts/:id/likes/p/:page/:items/t/:time/:newer
/posts/like/:id
/posts/liked/:id
/posts/delete/:id
/posts/undo/delete/:id
/posts/update
/posts/view/:id
/messages/threads/p/:page/:items/t/:time/:newer
/messages/:contact/timestamp
/messages/:contact/see
/attachments/all/p/:page/:items/t/:time/:newer
/attachments/u/:profile/p/:page/:items/t/:time/:newer
/attachments/s/:startup/p/:page/:items/t/:time/:newer
/attachments/s/:startup/r/:role/p/:page/:items/t/:time/:newer
/attachments/s/:startup/r/:role/a/:applicant/p/:page/:items/t/:time/:newer
/attachments/s/:startup/r/:role/g/:goal/p/:page/:items/t/:time/:newer
/messages/:contact/p/:page/:items/t/:time/:newer
/attachments/c/:channel/p/:page/:items/t/:time/:newer
/attachments/:id/comments/p/:page/:items/t/:time/:newer
````
# Skill

````
/skills/search/:keyword/p/:page/:items
````

# Notification

````
/notifications/p/:page/:items/t/:time/:newer
/notifications/p/:page/:items/t/:time/:newer/unread
/notifications/get/alerts
/notifications/see/all
/notifications/read/all
/notify/post/:post
/notify/startup/:id/invitations
/notify/startup/:id/invitations2
/notify/startup/:id/invitation/:invitation/accepted
/notify/startup/:id/invitation/:invitation/declined
/notify/startup/:id/submit
/notify/startup/:id/approve
/notify/startup/:id/reject
/notify/startup/:id/review
/notify/startup/:id/signed
/notify/startup/:id/publish
/notify/startup/:id/unpublish
/notify/startup/:id/development
/notify/startup/:id/launched
/notify/startup/:id/awarded
/notify/startup/:id/follow/by/:user
/notify/applicant/:applicant/applied
/notify/applicant/:applicant/approved
/notify/applicant/:applicant/rejected
/notify/applicant/:applicant/signed/all
/notify/applicant/:applicant/signed/startup
/notify/applicant/:applicant/signed/entrepreneur
/notify/applicant/:applicant/joined
/notify/applicant/:applicant/withdrawn
/notify/goal/:id/reached
/notify/goal/:id/approved
/notify/user/:user/new
/notify/user/:user/emails
/notify/user/:user/follower/:follower
````

# Interest

````
/interests/search/:keyword/p/:page/:items
````

# Language
````
/languages/search/:keyword/p/:page/:items
````

# Reports

````
/admin/reports/search
/admin/reports/status/:id
/admin/reports/approve/:id
/admin/reports/approved/:id
/admin/reports/decline/:id
/admin/reports/declined/:id
/admin/reports/search
/admin/reports/status/:id
/admin/reports/approve/:id
/admin/reports/approved/:id
/admin/reports/decline/:id
/admin/reports/declined/:id
````

# Logs

````
/admin/logs/search
````

# Frontend

````
/dashboard
/connect/linkedin/callback
/confirm-email
/reset-password
/logout
/legacy/profiles/:username
/legacy/profiles/:username/:tab
/people/follow/:id
/people/following/:id
/people/search
/people/admin/search
/people/:username
/people/:username/:tab
/people/:username/edit
/people/:username/edit/:tab
/users/profile/:userId
/settings
/settings/social-networks
/discover
/discover/startups
/discover/people
/discover/roles
/v1/startups/:id
/v1/startups/:id/:tab
/v1/create
/v1/startups/new
/v1/startups/new/:step
/v1/startups/new/:step/:tab
/v1/startups/edit/:id
/v1/startups/edit/:id/:step
/v1/startups/edit/:id/:step/:tab
/startups/:id
/startups/:id/:tab
/startups
/startups/drafts
/startups/active
/startups/applying
/startups/launched
/create
/startups/new
/startups/new/:tab
/startups/edit/:id
/startups/edit/:id/:tab
/startups/save
/startups/session/save
/startups/add
/startups/add/:tab
/startups/v/:id
/startups/v/:id/:tab
/startups/e/:id
/startups/e/:id/:tab
/startups/avatar-url
/startups/cover-url
/startups/submit/:id
/startups/discard/:id
/startups/restore/:id
/startups/unsubmit/:id
/startups/follow/:id
/startups/following/:id
/startups/approve/:id
/startups/reject/:id
/startups/review/:id
/startups/sign/:id
/startups/signed/:id
/startups/agreement/:id
/startups/publish/:id
/startups/unpublish/:id
/startups/develop/:id
/startups/launch/:id
/startups/award/:id
/startups/award/:id/update
/startups/award/:id/updated
/startups/block/:id
/startups/unblock/:id
/startups/invitation/:startupId
/startups/invitation/:memberId/:hash
/startups/invitation/:memberId/:hash/accept
/startups/invitation/:memberId/:hash/decline
/startups/list/:profile/p/:page/:items
/startups/fetch/:status/p/:page/:items
/startups/search
/startups/admin/search
/roles-pre-apply
/roles/:id
/roles/:id/:tab
/roles/:id/apply
/roles/:id/applicants/p/:page/:items/t/:time/:newer
/roles/search
/roles/admin/search
/roles
/roles/drafts
/roles/recruiting
/roles/applying
/roles/awarded
/roles/feed
/roles/feed/:format
/applications
/applications/running
/applications/won
/applications/lost
/applications/withdrawn
/applicants/:id
/applicants/:id/rate/:score
/applicants/:id/scores/p/:page/:items/t/:time/:newer
/applicants/:id/approve
/applicants/:id/reject
/applicants/:id/sign
/applicants/:id/signed
/applicants/:id/join
/applicants/:id/agreement
/applicants/:id/withdraw
/goals/:id
/goals/:id/reach
/goals/:id/approve
/posts/:id
/v1/messages
/v1/messages/user/:id
/v1/messages/compose
/v1/messages/threads/p/:page/:items/t/:time/:newer
/v1/messages/:contact/timestamp
/v1/messages/:contact/see
/attachments/all/p/:page/:items/t/:time/:newer
/attachments/u/:profile/p/:page/:items/t/:time/:newer
/attachments/s/:startup/p/:page/:items/t/:time/:newer
/attachments/s/:startup/r/:role/p/:page/:items/t/:time/:newer
/attachments/s/:startup/r/:role/a/:applicant/p/:page/:items/t/:time/:newer
/attachments/s/:startup/r/:role/g/:goal/p/:page/:items/t/:time/:newer
/messages/:contact/p/:page/:items/t/:time/:newer
/attachments/c/:channel/p/:page/:items/t/:time/:newer
/attachments/:id/comments/p/:page/:items/t/:time/:newer
/vcforu
/invest/:key
/dealrules
/messages
/messages/user/:id
messages/(.*)
messages/(\d+)
/users/search/:keyword/p/:page/:items
/users/search/offset
/users/:id
/users/list/:list/:profile/p/:page/:items
/skills/search/:keyword/p/:page/:items
/interests/search/:keyword/p/:page/:items
/languages/search/:keyword/p/:page/:items
/notifications/unread
/notifications/p/:page/:items/t/:time/:newer
/notifications/p/:page/:items/t/:time/:newer/unread
/notifications/get/alerts
/notifications/see/all
/notifications/read/all
/notify/post/:post
/notify/startup/:id/invitations
/notify/startup/:id/invitation/:invitation/accepted
/notify/startup/:id/invitation/:invitation/declined
/notify/startup/:id/submit
/notify/startup/:id/approve
/notify/startup/:id/reject
/notify/startup/:id/review
/notify/startup/:id/signed
/notify/startup/:id/publish
/notify/startup/:id/unpublish
/notify/startup/:id/development
/notify/startup/:id/launched
/notify/startup/:id/awarded
/notify/startup/:id/follow/by/:user
/notify/applicant/:applicant/applied
/notify/applicant/:applicant/approved
/notify/applicant/:applicant/rejected
/notify/applicant/:applicant/signed/all
/notify/applicant/:applicant/signed/startup
/notify/applicant/:applicant/signed/entrepreneur
/notify/applicant/:applicant/joined
/notify/applicant/:applicant/withdrawn
/notify/goal/:id/reached
/notify/goal/:id/approved
/notify/user/:user/new
/notify/user/:user/emails
/notify/user/:user/follower/:follower
/agreements/hellosign/callback
/cp
/cp/startups
/cp/people
/cp/roles
/cp/reports
/cp/logs
/cp/blacklist
/cp/loginas
/about
/press
/contact
/membership-agreement
/privacy-policy
/esa
/equity-sharing-agreement
/jscore
/schedule
/academy
/academy/:code
/tutorials/:code
/faq
/a/:code

````

