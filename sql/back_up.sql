PGDMP     2    '            	    z            personal_web    15.0    15.0     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                        1262    16398    personal_web    DATABASE     �   CREATE DATABASE personal_web WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_Indonesia.1252';
    DROP DATABASE personal_web;
                postgres    false            �            1259    16418    tb_projects    TABLE       CREATE TABLE public.tb_projects (
    id integer NOT NULL,
    name character varying NOT NULL,
    start_date date NOT NULL,
    end_date date NOT NULL,
    description text NOT NULL,
    technologies character varying[] NOT NULL,
    image character varying NOT NULL
);
    DROP TABLE public.tb_projects;
       public         heap    postgres    false            �            1259    16417    tb_blog_id_seq    SEQUENCE     �   CREATE SEQUENCE public.tb_blog_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 %   DROP SEQUENCE public.tb_blog_id_seq;
       public          postgres    false    217                       0    0    tb_blog_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public.tb_blog_id_seq OWNED BY public.tb_projects.id;
          public          postgres    false    216            h           2604    16421    tb_projects id    DEFAULT     l   ALTER TABLE ONLY public.tb_projects ALTER COLUMN id SET DEFAULT nextval('public.tb_blog_id_seq'::regclass);
 =   ALTER TABLE public.tb_projects ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    216    217    217            �          0    16418    tb_projects 
   TABLE DATA           g   COPY public.tb_projects (id, name, start_date, end_date, description, technologies, image) FROM stdin;
    public          postgres    false    217   g                  0    0    tb_blog_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('public.tb_blog_id_seq', 9, true);
          public          postgres    false    216            j           2606    16425    tb_projects tb_blog_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.tb_projects
    ADD CONSTRAINT tb_blog_pkey PRIMARY KEY (id);
 B   ALTER TABLE ONLY public.tb_projects DROP CONSTRAINT tb_blog_pkey;
       public            postgres    false    217            �   �   x���1O�0�����R��P�d��O�#;�lGj���Daa��Ó����
�}�VJ�j�������1O��5s���h�y�%�{ė�k�N�M��,;���yj�R��8NڻD�rF,A�`����D��h���#��Oэ(�*a!x7�ki�?_�Ӕ��h�dh��.5{�)AI3�`�S��Z:�VU�o�     