CREATE TABLE public.tb_device (
    id uuid NOT NULL,
    brand text,
    model text,
    operating_system text,
    operating_system_version text,
    timezone text,
    token text,
    unique_id text,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    users_ids uuid[]);

CREATE TABLE public.tb_group (
    id uuid NOT NULL,
    name text,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone);

CREATE TABLE public.tb_user (
    id uuid NOT NULL,
    email text,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    groups_ids uuid[]);

CREATE TABLE public.tb_user_addresses (
    parent_id uuid,
    id uuid NOT NULL,
    city text,
    country text,
    line_1 text,
    line_2 text,
    postal_code text,
    state text,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone);

CREATE TABLE public.tb_user_name (
    parent_id uuid,
    id uuid NOT NULL,
    first_name text,
    last_name text,
    middle_initial text,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone);

ALTER TABLE public.tb_device ADD CONSTRAINT chk_device__operating_system_version__empty_str CHECK (operating_system_version <> ''::TEXT);

ALTER TABLE public.tb_device ADD CONSTRAINT chk_device__token__empty_str CHECK (token <> ''::TEXT);

ALTER TABLE public.tb_device ADD CONSTRAINT chk_device__unique_id__empty_str CHECK (unique_id <> ''::TEXT);

ALTER TABLE public.tb_device ADD CONSTRAINT chk_device_brand_enum CHECK (brand = ANY (ARRAY['Apple'::text, 'Google'::text, 'Other'::text, 'Samsung'::text]));

ALTER TABLE public.tb_device ADD CONSTRAINT chk_device_model_enum CHECK (model = ANY (ARRAY['Android'::text, 'IPad'::text, 'IPhone'::text, 'Other'::text]));

ALTER TABLE public.tb_device ADD CONSTRAINT chk_device_operating_system_enum CHECK (operating_system = ANY (ARRAY['Android'::text, 'IOs'::text, 'Other'::text]));

ALTER TABLE public.tb_device ADD CONSTRAINT chk_device_timezone_enum CHECK (timezone = ANY (ARRAY['Berlin'::text, 'Central'::text, 'Eastern'::text, 'Kolkata'::text, 'London'::text, 'Mountain'::text, 'Pacific'::text, 'Paris'::text, 'Shanghai'::text, 'Sydney'::text, 'Tokyo'::text, 'Utc'::text]));

ALTER TABLE public.tb_device ADD CONSTRAINT pk_device_id PRIMARY KEY (id);

ALTER TABLE public.tb_group ADD CONSTRAINT chk_group__name__empty_str CHECK (name <> ''::TEXT);

ALTER TABLE public.tb_group ADD CONSTRAINT pk_group_id PRIMARY KEY (id);

ALTER TABLE public.tb_user ADD CONSTRAINT pk_user_id PRIMARY KEY (id);

ALTER TABLE public.tb_user_addresses ADD CONSTRAINT chk_addresses__city__empty_str CHECK (city <> ''::TEXT);

ALTER TABLE public.tb_user_addresses ADD CONSTRAINT chk_addresses__line_1__empty_str CHECK (line_1 <> ''::TEXT);

ALTER TABLE public.tb_user_addresses ADD CONSTRAINT chk_addresses__postal_code__empty_str CHECK (postal_code <> ''::TEXT);

ALTER TABLE public.tb_user_addresses ADD CONSTRAINT chk_addresses_country_enum CHECK (country = ANY (ARRAY['Afghanistan'::text, 'Albania'::text, 'Algeria'::text, 'Andorra'::text, 'Angola'::text, 'AntiguaAndBarbuda'::text, 'Argentina'::text, 'Armenia'::text, 'Australia'::text, 'Austria'::text, 'Azerbaijan'::text, 'Bahamas'::text, 'Bahrain'::text, 'Bangladesh'::text, 'Barbados'::text, 'Belarus'::text, 'Belgium'::text, 'Belize'::text, 'Benin'::text, 'Bhutan'::text, 'Bolivia'::text, 'BosniaAndHerzegovina'::text, 'Botswana'::text, 'Brazil'::text, 'Brunei'::text, 'Bulgaria'::text, 'BurkinaFaso'::text, 'Burundi'::text, 'CaboVerde'::text, 'Cambodia'::text, 'Cameroon'::text, 'Canada'::text, 'CentralAfricanRepublic'::text, 'Chad'::text, 'Chile'::text, 'China'::text, 'Colombia'::text, 'Comoros'::text, 'Congo'::text, 'CostaRica'::text, 'Croatia'::text, 'Cuba'::text, 'Cyprus'::text, 'Czechia'::text, 'DemocraticRepublicOfCongo'::text, 'Denmark'::text, 'Djibouti'::text, 'Dominica'::text, 'DominicanRepublic'::text, 'Ecuador'::text, 'Egypt'::text, 'ElSalvador'::text, 'EquatorialGuinea'::text, 'Eritrea'::text, 'Estonia'::text, 'Eswatini'::text, 'Ethiopia'::text, 'Fiji'::text, 'Finland'::text, 'France'::text, 'Gabon'::text, 'Gambia'::text, 'Georgia'::text, 'Germany'::text, 'Ghana'::text, 'Greece'::text, 'Grenada'::text, 'Guatemala'::text, 'Guinea'::text, 'GuineaBissau'::text, 'Guyana'::text, 'Haiti'::text, 'Honduras'::text, 'Hungary'::text, 'Iceland'::text, 'India'::text, 'Indonesia'::text, 'Iran'::text, 'Iraq'::text, 'Ireland'::text, 'Israel'::text, 'Italy'::text, 'IvoryCoast'::text, 'Jamaica'::text, 'Japan'::text, 'Jordan'::text, 'Kazakhstan'::text, 'Kenya'::text, 'Kiribati'::text, 'Kosovo'::text, 'Kuwait'::text, 'Kyrgyzstan'::text, 'Laos'::text, 'Latvia'::text, 'Lebanon'::text, 'Lesotho'::text, 'Liberia'::text, 'Libya'::text, 'Liechtenstein'::text, 'Lithuania'::text, 'Luxembourg'::text, 'Madagascar'::text, 'Malawi'::text, 'Malaysia'::text, 'Maldives'::text, 'Mali'::text, 'Malta'::text, 'MarshallIslands'::text, 'Mauritania'::text, 'Mauritius'::text, 'Mexico'::text, 'Micronesia'::text, 'Moldova'::text, 'Monaco'::text, 'Mongolia'::text, 'Montenegro'::text, 'Morocco'::text, 'Mozambique'::text, 'Myanmar'::text, 'Namibia'::text, 'Nauru'::text, 'Nepal'::text, 'Netherlands'::text, 'NewZealand'::text, 'Nicaragua'::text, 'Niger'::text, 'Nigeria'::text, 'NorthKorea'::text, 'NorthMacedonia'::text, 'Norway'::text, 'Oman'::text, 'Pakistan'::text, 'Palau'::text, 'Palestine'::text, 'Panama'::text, 'PapuaNewGuinea'::text, 'Paraguay'::text, 'Peru'::text, 'Philippines'::text, 'Poland'::text, 'Portugal'::text, 'Qatar'::text, 'Romania'::text, 'Russia'::text, 'Rwanda'::text, 'SaintKittsAndNevis'::text, 'SaintLucia'::text, 'SaintVincentAndTheGrenadines'::text, 'Samoa'::text, 'SanMarino'::text, 'SaoTomeAndPrincipe'::text, 'SaudiArabia'::text, 'Senegal'::text, 'Serbia'::text, 'Seychelles'::text, 'SierraLeone'::text, 'Singapore'::text, 'Slovakia'::text, 'Slovenia'::text, 'SolomonIslands'::text, 'Somalia'::text, 'SouthAfrica'::text, 'SouthKorea'::text, 'SouthSudan'::text, 'Spain'::text, 'SriLanka'::text, 'Sudan'::text, 'Suriname'::text, 'Sweden'::text, 'Switzerland'::text, 'Syria'::text, 'Taiwan'::text, 'Tajikistan'::text, 'Tanzania'::text, 'Thailand'::text, 'TimorLeste'::text, 'Togo'::text, 'Tonga'::text, 'TrinidadAndTobago'::text, 'Tunisia'::text, 'Turkey'::text, 'Turkmenistan'::text, 'Tuvalu'::text, 'Uganda'::text, 'Ukraine'::text, 'UnitedArabEmirates'::text, 'UnitedKingdom'::text, 'UnitedStatesOfAmerica'::text, 'Uruguay'::text, 'Uzbekistan'::text, 'Vanuatu'::text, 'VaticanCity'::text, 'Venezuela'::text, 'Vietnam'::text, 'Yemen'::text, 'Zambia'::text, 'Zimbabwe'::text]));

ALTER TABLE public.tb_user_addresses ADD CONSTRAINT chk_addresses_state_enum CHECK (state = ANY (ARRAY['Alabama'::text, 'Alaska'::text, 'AmericanSamoa'::text, 'Arizona'::text, 'Arkansas'::text, 'California'::text, 'Colorado'::text, 'Connecticut'::text, 'Delaware'::text, 'DistrictOfColumbia'::text, 'Florida'::text, 'Georgia'::text, 'Guam'::text, 'Hawaii'::text, 'Idaho'::text, 'Illinois'::text, 'Indiana'::text, 'Iowa'::text, 'Kansas'::text, 'Kentucky'::text, 'Louisiana'::text, 'Maine'::text, 'Maryland'::text, 'Massachusetts'::text, 'Michigan'::text, 'Minnesota'::text, 'Mississippi'::text, 'Missouri'::text, 'Montana'::text, 'Nebraska'::text, 'Nevada'::text, 'NewHampshire'::text, 'NewJersey'::text, 'NewMexico'::text, 'NewYork'::text, 'NorthCarolina'::text, 'NorthDakota'::text, 'NorthernMarianaIslands'::text, 'Ohio'::text, 'Oklahoma'::text, 'Oregon'::text, 'Pennsylvania'::text, 'PuertoRico'::text, 'RhodeIsland'::text, 'SouthCarolina'::text, 'SouthDakota'::text, 'Tennessee'::text, 'Texas'::text, 'Utah'::text, 'Vermont'::text, 'VirginIslands'::text, 'Virginia'::text, 'Washington'::text, 'WestVirginia'::text, 'Wisconsin'::text, 'Wyoming'::text]));

ALTER TABLE public.tb_user_addresses ADD CONSTRAINT pk_user_addresses_id PRIMARY KEY (id);

ALTER TABLE public.tb_user_name ADD CONSTRAINT chk_name__first_name__empty_str CHECK (first_name <> ''::TEXT);

ALTER TABLE public.tb_user_name ADD CONSTRAINT chk_name__last_name__empty_str CHECK (last_name <> ''::TEXT);

ALTER TABLE public.tb_user_name ADD CONSTRAINT pk_user_name_id PRIMARY KEY (id);

ALTER TABLE public.tb_user_addresses ADD CONSTRAINT fk_user_addresses__parent_id__user FOREIGN KEY (parent_id) REFERENCES public.tb_user (id) ON DELETE CASCADE;

ALTER TABLE public.tb_user_name ADD CONSTRAINT fk_user_name__parent_id__user FOREIGN KEY (parent_id) REFERENCES public.tb_user (id) ON DELETE CASCADE;

DROP EXTENSION plpgsql;

