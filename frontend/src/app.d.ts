import { SupabaseClient, Session, User } from "@supabase/supabase-js";

// See https://svelte.dev/docs/kit/types#app.d.ts
// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		interface Locals {
			supabase: SupabaseClient;
			getUserAndSession(): Promise<{ user: User | null; session: Session | null }>;
			getSession(): Promise<Session | null>;
			session: Session | null;
			getUser(): Promise<User | null>;
			user: User | null;
		}
		interface PageData {
			session: Session | null;
		}
		// interface PageState {}
		// interface Platform {}
	}
}

export {};
