import SuperTokens from 'supertokens-web-js';
import Session from 'supertokens-web-js/recipe/session';
import ThirdParty from 'supertokens-web-js/recipe/thirdparty'
import EmailPassword from 'supertokens-web-js/recipe/emailpassword'

export default function InitializeSuperTokens() {
    SuperTokens.init({
        appInfo: {
            apiDomain: "http://localhost:8081",
            apiBasePath: "/auth",
            appName: "...",
        },
        recipeList: [
            Session.init(),
            EmailPassword.init(),
            ThirdParty.init()
        ],
    });    
}