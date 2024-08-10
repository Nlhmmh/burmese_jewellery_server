INSERT INTO account_admins (account_admin_id, mail, password, account_admin_role, account_admin_status, created_at, updated_at)
VALUES
  ('b106e919-ab73-4444-a86c-4b38ea22a0fe', 'admin@gmail.com', '$2a$10$On3UdcaBCZNYOwugRB9zuu5/AV8ywYVboiDqggpe0eaj8xK0Z1eTK', 'admin', 'active', NOW(), NOW()),
  ('b106e919-ab73-4444-a86c-4b38ea22a0f1', 'staff@gmail.com', '$2a$10$On3UdcaBCZNYOwugRB9zuu5/AV8ywYVboiDqggpe0eaj8xK0Z1eTK', 'staff', 'active', NOW() + INTERVAL '1 second', NOW() + INTERVAL '1 second')
;

INSERT INTO categories (category_id, name, description, image_url)
VALUES
  ('b106e919-ab73-48f3-a86c-4b38ea22a0fe', 'Engagement Rings', 'Engagement Rings', 'http://localhost:8077/api/file/engagement-rings.jpg'),
  ('b106e919-ab73-48f3-a86c-4b38ea22a0ff', 'Wedding Rings', 'Wedding Rings', 'http://localhost:8077/api/file/wedding-rings.jpg'),
  ('b106e919-ab73-48f3-a86c-4b38ea22a0f1', 'Necklace', 'Necklace', 'http://localhost:8077/api/file/necklace.jpg'),
  ('b106e919-ab73-48f3-a86c-4b38ea22a0f2', 'Earrings', 'Earrings', 'http://localhost:8077/api/file/earrings.jpg'),
  ('b106e919-ab73-48f3-a86c-4b38ea22a0f3', 'Bracelets', 'Bracelets', 'http://localhost:8077/api/file/bracelets.jpg')
;

INSERT INTO gems (gem_id, name)
VALUES
  ('46e78b04-955d-4f0c-a751-dd1dcb2d7248', 'Diamond'),
  ('46e78b04-955d-4f0c-a751-dd1dcb2d7242', 'Ruby'),
  ('46e78b04-955d-4f0c-a751-dd1dcb2d7243', 'Sapphire'),
  ('46e78b04-955d-4f0c-a751-dd1dcb2d7244', 'Emerald'),
  ('46e78b04-955d-4f0c-a751-dd1dcb2d7245', 'Jade')
;

INSERT INTO materials (material_id, name)
VALUES
  ('c0979115-2bb3-4c50-a29d-64f3b8784a37', 'Gold'),
  ('c0979115-2bb3-4c50-a29d-64f3b8784a32', 'Platinum'),
  ('c0979115-2bb3-4c50-a29d-64f3b8784a33', 'Silver')
;

INSERT INTO accounts (account_id, login_type, mail, password, phone, account_status)
VALUES
  ('b106e919-ab73-48f3-a86c-4b38ea22a888', 'email', 'testuser@gmail.com', '$2a$10$On3UdcaBCZNYOwugRB9zuu5/AV8ywYVboiDqggpe0eaj8xK0Z1eTK', '08077779999', 'active')
;

INSERT INTO account_profiles (account_id, first_name, last_name, birthday, gender)
VALUES
  ('b106e919-ab73-48f3-a86c-4b38ea22a888', 'John', 'Tester', '1990-11-01', 'male')
;

INSERT INTO faqs (faq_id, question, answer, is_active, created_at, updated_at)
VALUES
  ('c0979115-2bb3-4c50-a29d-64f3b8784888', 'How can I make returns or exchanges?', 'Products are accepted for credit or exchange in the U.S. within 30 days of purchase. ', true, NOW(), NOW()),
  ('c0979115-2bb3-4c50-a29d-64f3b8784887', 'How do I return an online purchase?', 'Simply visit this page to print a return label.', false, NOW() + INTERVAL '1 second', NOW() + INTERVAL '1 second')
;

INSERT INTO jewelleries (jewellery_id, category_id, gem_id, material_id, name, description, price, image_url, is_published, quantity)
VALUES
  (
    'c0979115-2bb3-9999-a29d-64f3b8784888',
    'b106e919-ab73-48f3-a86c-4b38ea22a0ff',
    '46e78b04-955d-4f0c-a751-dd1dcb2d7248',
    'c0979115-2bb3-4c50-a29d-64f3b8784a37',
    'Wedding Forever',
    'Wedding Forever in Platinum with a Half-circle of Diamonds, 2.2 mm',
    '100',
    'http://localhost:8077/api/file/gold-wedding-rings-pairs-1.jpg',
    true,
    100
  ),
  (
    'c0979115-2bb3-9999-a29d-64f3b8784811',
    'b106e919-ab73-48f3-a86c-4b38ea22a0fe',
    '46e78b04-955d-4f0c-a751-dd1dcb2d7248',
    'c0979115-2bb3-4c50-a29d-64f3b8784a33',
    'Lasting Engagement',
    'Wire Engagment Ring in Yellow Gold, 2.2 mm',
    '100',
    'http://localhost:8077/api/file/diamond-wedding-rings-paairs-1.jpg',
    true,
    100
  ),
  (
    'c0979115-2bb3-9999-a29d-64f3b8784812',
    'b106e919-ab73-48f3-a86c-4b38ea22a0f1',
    '46e78b04-955d-4f0c-a751-dd1dcb2d7248',
    'c0979115-2bb3-4c50-a29d-64f3b8784a37',
    'Necklace',
    'Necklace in Yellow Gold',
    '100',
    'http://localhost:8077/api/file/necklace.jpg',
    true,
    100
  ),
  (
    'c0979115-2bb3-9999-a29d-64f3b8784813',
    'b106e919-ab73-48f3-a86c-4b38ea22a0f2',
    '46e78b04-955d-4f0c-a751-dd1dcb2d7248',
    'c0979115-2bb3-4c50-a29d-64f3b8784a37',
    'Lock Small Earrings',
    'Lock Small Earrings in Yellow Gold',
    '100',
    'http://localhost:8077/api/file/gold-earrings-1.jpg',
    true,
    100
  ),
  (
    'c0979115-2bb3-9999-a29d-64f3b8784814',
    'b106e919-ab73-48f3-a86c-4b38ea22a0f3',
    '46e78b04-955d-4f0c-a751-dd1dcb2d7248',
    'c0979115-2bb3-4c50-a29d-64f3b8784a37',
    'T Smile Pendant',
    'T Smile Pendant in Yellow Gold',
    '100',
    'http://localhost:8077/api/file/gold-bracelet-1.jpg',
    true,
    100
  ),
  (
    'c0979115-2bb3-9999-a29d-64f3b8784815',
    'b106e919-ab73-48f3-a86c-4b38ea22a0ff',
    '46e78b04-955d-4f0c-a751-dd1dcb2d7248',
    'c0979115-2bb3-4c50-a29d-64f3b8784a37',
    'Wedding Forever 1',
    'Wedding Forever in Platinum with a Half-circle of Diamonds, 2.2 mm',
    '100',
    'http://localhost:8077/api/file/gold-wedding-rings-pairs-1.jpg',
    true,
    100
  ),
  (
    'c0979115-2bb3-9999-a29d-64f3b8784816',
    'b106e919-ab73-48f3-a86c-4b38ea22a0fe',
    '46e78b04-955d-4f0c-a751-dd1dcb2d7248',
    'c0979115-2bb3-4c50-a29d-64f3b8784a33',
    'Lasting Engagement 1',
    'Wire Engagment Ring in Yellow Gold, 2.2 mm',
    '100',
    'http://localhost:8077/api/file/diamond-wedding-rings-paairs-1.jpg',
    true,
    100
  ),
  (
    'c0979115-2bb3-9999-a29d-64f3b8784817',
    'b106e919-ab73-48f3-a86c-4b38ea22a0f1',
    '46e78b04-955d-4f0c-a751-dd1dcb2d7248',
    'c0979115-2bb3-4c50-a29d-64f3b8784a37',
    'Necklace 1',
    'Necklace in Yellow Gold',
    '100',
    'http://localhost:8077/api/file/necklace.jpg',
    true,
    100
  ),
  (
    'c0979115-2bb3-9999-a29d-64f3b8784818',
    'b106e919-ab73-48f3-a86c-4b38ea22a0f2',
    '46e78b04-955d-4f0c-a751-dd1dcb2d7248',
    'c0979115-2bb3-4c50-a29d-64f3b8784a37',
    'Lock Small Earrings 1',
    'Lock Small Earrings in Yellow Gold',
    '100',
    'http://localhost:8077/api/file/gold-earrings-1.jpg',
    true,
    100
  ),
  (
    'c0979115-2bb3-9999-a29d-64f3b8784819',
    'b106e919-ab73-48f3-a86c-4b38ea22a0f3',
    '46e78b04-955d-4f0c-a751-dd1dcb2d7248',
    'c0979115-2bb3-4c50-a29d-64f3b8784a37',
    'T Smile Pendant 1',
    'T Smile Pendant in Yellow Gold',
    '100',
    'http://localhost:8077/api/file/gold-bracelet-1.jpg',
    true,
    100
  )
;
