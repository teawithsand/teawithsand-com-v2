export const applyIntlErrorHack = () => {
    // TODO(teawithsand): disable this in prod builds
    // eslint-disable-next-line
    const consoleError = console.error.bind(console);
    // eslint-disable-next-line
    console.error = (candidate, ...args) => {
        const expr = /@formatjs\/intl Error MISSING_TRANSLATION/
        if (
            (typeof candidate === 'string' && expr.test(candidate)) ||
            (candidate instanceof Error && typeof candidate.message === "string" && expr.test(candidate.message))
        ) {
            return;
        }
        consoleError(candidate, ...args);
    }
}