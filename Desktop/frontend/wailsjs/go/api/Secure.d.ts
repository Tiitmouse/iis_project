// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {api} from '../models';

export function CreateContact(arg1:api.Contact):Promise<api.Contact>;

export function DeleteContact(arg1:string):Promise<void>;

export function FetchContact(arg1:string):Promise<api.Contact>;

export function FetchContacts():Promise<Array<api.Contact>>;

export function Login(arg1:string,arg2:string):Promise<void>;

export function Logout():Promise<void>;

export function UpdateContact(arg1:string,arg2:api.Contact):Promise<api.Contact>;
