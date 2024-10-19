// now this I DO remember, if hazily
// it was written in response to someone on TC-39 asking for feedback or use cases for decorators on twitter
// it was right after classpass had laid me off and I was working on plate scraper again and struggling
// with input validation for the endpoints and wishing there was something like the kotlin/java
// decorators I had seen in the new partners-service kotlin backend

// I think this is mostly pseudocode and I had no idea what the fuck I was talking about, so it might not make much sense

// I should say up front I have no idea if this is actually possible in JS/TS with or without the decorators proposal. I've seen a couple things that seem kind of like it (which I'll link), but nothing that actually is it. The idea honestly isn't even 100% formed in my head.

// I read the (updated) proposal and it seems like it might be possible to use a method decorator like @logged with a factory function to validate the types of params at runtime. The decorator factory would take a param to test and a validation function, then run the method if the validation succeeds, or take some other action (like throwing a user defined error) if the validation fails.

// I think this could be particularly useful for typescript, since it could let you bridge the gap between compile time and runtime type safety. Ordinarily, the compiler can't know the types of fields on req.body/params/query in advance because they're supplied by the client, which is outside the program's control. But I can imagine a scenario where if the decorators validate a field, the function would know its type and be able to use it in a typesafe way.

// "lines and lines of manual typechecking" version

router.post(`/api/recipe`, async (req, res, next) => {
  try {
      const { text, title, tags } = req.body;
      if (
        typeof text === 'string' &&
        typeof title === 'string' &&
        Array.isArray(tags) &&
        tags.every(tag => typeof tag === 'string')
      ) {
        // adding new recipe to a db
        const recipe = await recipeRepository.insert({
          text,
          title,
          // get the tag if it already exists, insert it and return it if not for all the tags in the array
          tags: await tagRepository.getOrInsert(tags),
        });
        // your db might refuse to take the wrong types and throw an error, but I can imagine situations where you'd want to avoid the call if you know it will fail, or other situations that require typechecking
        res.json(recipe);
      } else {
        // this could certainly be much more verbose/more lines if you wanted to specify which field had an invalid type
        throw new Error("recipe has invalid types")
      }
  } catch (error) {
    next(error);
  }
});

// for a regular function
const add = (a, b) => {
  if (typeof a === 'number' && typeof b === 'number') {
    return a + b
  } else {
    throw new Error(":(")
  }
}

// thing I hope might be possible with decorators

// each decorator would validate the type of the argument and take some action (like throwing an error) if it failed (or gather up all the failures at the end and throw)
// if all passed, the main function would run and know the types of each param (in TS) or just run (in JS)
// Even if decorators can only be used on class methods, you could make your endpoint handlers class methods, then apply the decorators

const bodyParam = (bodyParamName, validationFn) => function(f) {
  const name = f.name;
  function wrapped(req, res, next) {
    if (validationFn(req.body[bodyParamName])){
      f.call(this, req, res, next)
    } else {
      throw new Error(`${bodyParamName} of invalid type: ${typeof req.body[bodyParamName]}`)
    }
  }
  Object.defineProperty(wrapped, "name", { value: name, configurable: true });
  return wrapped;
}

const isStr = value => typeof value === 'string'
const isArrOfStr = value => Array.isArray === value && value.every(isStr)

// I may be misunderstanding the evaluation -> call -> apply process; would you have to create the decorators in advance using the bodyParam factory, or is bodyParam() just referentially tansparent (ie, the return value (the actual decorator) would be used)?
@bodyParam("text", isStr)
@bodyParam("title", isStr)
@bodyParam("tags", isArrOfStr)
router.post(`/api/recipe`, async (req, res, next) => {
  try {
    const { text, title, tags } = req.body;
    const recipe = await recipeRepository.insert({
      text,
      title,
      tags: await tagRepository.getOrInsert(tags),
    });
    res.json(recipe);
  } catch (error) {
    next(error);
  }
});

// regular function
@param(a, isNum)
@param(b, isNum)
const add = (a, b) => a + b

// You could also write one that would try and coerce the value into the correct type and pass if it's valid, or fail if it's not. For instance, req.params.id on a route like
// /api/user/:id will always be a string, but you probably want it as a number
// so @routeParam("id", isNum)
// would pass /api/user/3
// would fail /api/user/yellow
// Other kinds of transforms might be possible too (transform all the tags from strings to lowercase)

// I could also imagine a typescript version that does something like @param(paramName): type
// perhaps devs (or a library) could create decorator factories with partial application like this
const bodyParamStr = bodyParamName => bodyParam(bodyParamName, isStr)
const bodyParamArrOfStr = bodyParamName => bodyParam(bodyParamName, isArrOfStr)

bodyParamStr("text")
bodyParamStr("title")
bodyParamArrOfStr("tags")
router.post(`/api/recipe`, myEndpointHandler)

// "Prior art"

// this is basically a ripoff of something that's doable and quite handy in (statically typed, totally different JVM language) Kotlin
/*
@GET
@Path("/venues/{venue_id}/schedules/{schedule_id}/roster")
fun getRoster(
    // these bits right here
    @PathParam("venue_id") @UnwrapValidatedValue @NotNull venueId: IntParam,
    @PathParam("schedule_id") @UnwrapValidatedValue @NotNull scheduleId: IntParam,
    @QueryParam("include_emails_override") includeEmailsOverride: Boolean?
): ReservationsResponse {
    return ReservationsResponse(
        controller.getRoster(venueId.get(), scheduleId.get(), includeEmailsOverride ?: false)
    )
}
*/

// I've seen a couple things floating around (mostly for typescript) that are kind of LIKE this
// https://github.com/MichalLytek/class-transformer-validator, which uses https://github.com/typestack/class-validator
// It makes use of the experimental decorator and reflect metadata settings in typescript described here https://www.typescriptlang.org/docs/handbook/decorators.html and https://www.npmjs.com/package/reflect-metadata
// this person had a request for similar capabilities with TS and express; no one could help them https://stackoverflow.com/questions/54063704/add-properties-to-the-req-object-in-expressjs-with-typescript

// I hope this is helpful and not way too much or too little information. Please let me know if there's anything I can clarify or expand on
