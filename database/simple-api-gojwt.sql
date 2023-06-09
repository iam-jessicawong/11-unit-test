PGDMP     2    +                {            simple-api-gojwt    15.2    15.2     	           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            
           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    16451    simple-api-gojwt    DATABASE     �   CREATE DATABASE "simple-api-gojwt" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_Indonesia.1252';
 "   DROP DATABASE "simple-api-gojwt";
                postgres    false            �            1259    16462    products    TABLE     �   CREATE TABLE public.products (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    title character varying(100) NOT NULL,
    description text NOT NULL,
    user_id bigint
);
    DROP TABLE public.products;
       public         heap    postgres    false            �            1259    16461    products_id_seq    SEQUENCE     x   CREATE SEQUENCE public.products_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.products_id_seq;
       public          postgres    false    217                       0    0    products_id_seq    SEQUENCE OWNED BY     C   ALTER SEQUENCE public.products_id_seq OWNED BY public.products.id;
          public          postgres    false    216            �            1259    16453    users    TABLE     H  CREATE TABLE public.users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    full_name character varying(100) NOT NULL,
    email character varying(100) NOT NULL,
    password text NOT NULL,
    role character varying(5) DEFAULT 'user'::character varying NOT NULL
);
    DROP TABLE public.users;
       public         heap    postgres    false            �            1259    16452    users_id_seq    SEQUENCE     u   CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.users_id_seq;
       public          postgres    false    215                       0    0    users_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
          public          postgres    false    214            l           2604    16465    products id    DEFAULT     j   ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_id_seq'::regclass);
 :   ALTER TABLE public.products ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    217    216    217            j           2604    16456    users id    DEFAULT     d   ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);
 7   ALTER TABLE public.users ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    215    214    215                      0    16462    products 
   TABLE DATA           [   COPY public.products (id, created_at, updated_at, title, description, user_id) FROM stdin;
    public          postgres    false    217   ^                 0    16453    users 
   TABLE DATA           ]   COPY public.users (id, created_at, updated_at, full_name, email, password, role) FROM stdin;
    public          postgres    false    215   1                  0    0    products_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public.products_id_seq', 5, true);
          public          postgres    false    216                       0    0    users_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.users_id_seq', 7, true);
          public          postgres    false    214            n           2606    16488    users idx_users_email 
   CONSTRAINT     Q   ALTER TABLE ONLY public.users
    ADD CONSTRAINT idx_users_email UNIQUE (email);
 ?   ALTER TABLE ONLY public.users DROP CONSTRAINT idx_users_email;
       public            postgres    false    215            r           2606    16469    products products_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.products DROP CONSTRAINT products_pkey;
       public            postgres    false    217            p           2606    16460    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    215            s           2606    16470    products fk_users_product    FK CONSTRAINT     �   ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_users_product FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE SET NULL;
 C   ALTER TABLE ONLY public.products DROP CONSTRAINT fk_users_product;
       public          postgres    false    3184    215    217            t           2606    16496    products fk_users_products    FK CONSTRAINT     �   ALTER TABLE ONLY public.products
    ADD CONSTRAINT fk_users_products FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE SET NULL;
 D   ALTER TABLE ONLY public.products DROP CONSTRAINT fk_users_products;
       public          postgres    false    3184    215    217               �   x����n�0Fg�)�S"���cW�t��8�J���D+�4R7��Ƞ�"m-o���upX��X�����&��Ә_�=��"7�Xz�|���r�f>�1�V��Ų���n�C��]ڽ����y��K徖��1�@>��ǀ�/R+�p-S����Y����{�D����hG�Y��C9�s�2���Jk� F.a�         *  x����n�@�5<�w���q,+�8c�Q�*%� C��B�Oߦ&��&MN�ɟq>�hP�=�� � j d
�ABt�?�4��$:[Y����q$e�t��]���`*ܝ	�6u�j
�ɦcYא���圏�G"=��#{P�'��u 6�nh �0���,7�̢w��9�kՁl���I�z�s_,&�^����".y�c>��lg����+MU�~��{�ҰF�T7�YEJ�����뻏�֛��|M�KvjL���+�y�vKʏ�eJ7_>��u�<�^���_	#|D     