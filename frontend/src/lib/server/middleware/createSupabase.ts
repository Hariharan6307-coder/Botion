import { PUBLIC_SUPABASE_URL, PUBLIC_SUPABASE_ANON_KEY } from "$env/static/public";
import { createServerClient } from "@supabase/ssr";
import type { Handle } from "@sveltejs/kit";

const createSupabase:Handle = async ({ event, resolve }) => {
    event.locals.supabase = createServerClient(PUBLIC_SUPABASE_URL, PUBLIC_SUPABASE_ANON_KEY, {
        cookies: {
            getAll() {
                return event.cookies.getAll();
            },
            setAll(cookiesToSet) {
                cookiesToSet.forEach(( { name, value, options }) => {
                    event.cookies.set(name, value, { ...options, path: '/'})
                });
            }
        }
    });

    event.locals.getUserAndSession = async () => {
        const { data: { user }, error } = await event.locals.supabase.auth.getUser();
        if (error) {
            return { user: null, session: null };
        }
        const { data: { session } } = await event.locals.supabase.auth.getSession();
        if (!session) {
            return { user: null, session: null };
        }
        return { user, session };
    }

    event.locals.getSession = async () => {
        const { session } = await event.locals.getUserAndSession();
        return session;
    }
    event.locals.session = await event.locals.getSession();

    event.locals.getUser = async () => {
        const { user } = await event.locals.getUserAndSession();
        return user;
    }
    event.locals.user = await event.locals.getUser();

    return resolve(event, {
        filterSerializedResponseHeaders(name) {
            return name == 'content-range' || name == 'x-supabase-api-version';
        }
    })
}

export default createSupabase;